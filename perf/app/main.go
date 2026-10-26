// Command goredis-bench is a small bench app that drives go-redis with
// GET/SET/HGET/HSET/HGETALL, with or without autopipelining.
//
// Settings come from scenario.params in $BENCH_CONFIG:
//
//	autopipeline   bool  use client.AutoPipeline() instead of the plain client (default false)
//	concurrency    int   goroutines issuing commands (default 64)
//	key_space      int   number of distinct keys (default 1000)
//	payload_bytes  int   value size for string commands and hash fields (default 64)
//	hash_fields    int   fields per hash (default 10)
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

type config struct {
	Scenario struct {
		Name   string         `json:"name"`
		Params map[string]any `json:"params"`
	} `json:"scenario"`
	Redis struct {
		URL    string `json:"url"`
		Prefix string `json:"prefix"`
	} `json:"redis"`
	Run struct {
		WarmupS   float64 `json:"warmup_s"`
		DurationS float64 `json:"duration_s"`
	} `json:"run"`
}

type settings struct {
	autopipeline bool
	concurrency  int
	keySpace     int
	payloadBytes int
	hashFields   int
}

// cmdable is the subset both *redis.Client and *redis.AutoPipeliner satisfy.
type cmdable interface {
	Set(ctx context.Context, key string, value any, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	HSet(ctx context.Context, key string, values ...any) *redis.IntCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "goredis:", err)
		os.Exit(1)
	}
}

func run() error {
	raw, err := os.ReadFile(os.Getenv("BENCH_CONFIG"))
	if err != nil {
		return fmt.Errorf("read BENCH_CONFIG: %w", err)
	}
	var cfg config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return fmt.Errorf("parse BENCH_CONFIG: %w", err)
	}
	s, err := parseSettings(cfg.Scenario.Params)
	if err != nil {
		return err
	}
	scenario := cfg.Scenario.Name
	switch scenario {
	case "set", "get", "hset", "hget", "hgetall":
	default:
		return fmt.Errorf("unknown scenario %q", scenario)
	}

	opts, err := redis.ParseURL(cfg.Redis.URL)
	if err != nil {
		return fmt.Errorf("parse redis url: %w", err)
	}
	opts.PoolSize = s.concurrency
	client := redis.NewClient(opts)
	defer client.Close()

	var target cmdable = client
	if s.autopipeline {
		ap, err := client.AutoPipeline()
		if err != nil {
			return fmt.Errorf("autopipeline: %w", err)
		}
		defer ap.Close()
		target = ap
	}
	fmt.Fprintf(os.Stderr, "goredis: scenario=%s autopipeline=%t concurrency=%d prefix=%s\n",
		scenario, s.autopipeline, s.concurrency, cfg.Redis.Prefix)

	ctx := context.Background()
	keys := make([]string, s.keySpace)
	for i := range keys {
		keys[i] = cfg.Redis.Prefix + "k:" + strconv.Itoa(i)
	}
	defer func() {
		// Only our own prefixed keys, in chunks.
		for i := 0; i < len(keys); i += 500 {
			end := min(i+500, len(keys))
			client.Del(ctx, keys[i:end]...)
		}
	}()

	payload := make([]byte, s.payloadBytes)
	for i := range payload {
		payload[i] = 'x'
	}
	value := string(payload)
	fields := make([]string, s.hashFields)
	hashArgs := make([]any, 0, 2*s.hashFields)
	for i := range fields {
		fields[i] = "f" + strconv.Itoa(i)
		hashArgs = append(hashArgs, fields[i], value)
	}

	if err := seed(ctx, client, scenario, keys, value, hashArgs); err != nil {
		return fmt.Errorf("seed: %w", err)
	}

	op := func(rng *rand.Rand) error {
		key := keys[rng.Intn(len(keys))]
		switch scenario {
		case "set":
			return target.Set(ctx, key, value, 0).Err()
		case "get":
			return target.Get(ctx, key).Err()
		case "hset":
			return target.HSet(ctx, key, fields[rng.Intn(len(fields))], value).Err()
		case "hget":
			return target.HGet(ctx, key, fields[rng.Intn(len(fields))]).Err()
		default:
			return target.HGetAll(ctx, key).Err()
		}
	}

	if cfg.Run.WarmupS > 0 {
		drive(s.concurrency, time.Duration(cfg.Run.WarmupS*float64(time.Second)), op)
	}
	res := drive(s.concurrency, time.Duration(cfg.Run.DurationS*float64(time.Second)), op)

	sort.Slice(res.latencies, func(i, j int) bool { return res.latencies[i] < res.latencies[j] })
	pct := func(q int) float64 {
		n := len(res.latencies)
		if n == 0 {
			return 0
		}
		return float64(res.latencies[min(n-1, n*q/100)]) / 1e6
	}
	ops := float64(len(res.latencies)) / res.elapsed.Seconds()

	out := map[string]any{"metrics": map[string]any{
		"throughput_ops": map[string]any{"value": ops, "unit": "ops/s", "better": "higher"},
		"latency_p50_ms": map[string]any{"value": pct(50), "unit": "ms", "better": "lower"},
		"latency_p95_ms": map[string]any{"value": pct(95), "unit": "ms", "better": "lower"},
		"latency_p99_ms": map[string]any{"value": pct(99), "unit": "ms", "better": "lower"},
		"errors":         map[string]any{"value": res.errors, "unit": "count", "better": "lower"},
	}}
	b, err := json.Marshal(out)
	if err != nil {
		return err
	}
	return os.WriteFile(os.Getenv("BENCH_RESULT"), b, 0o644)
}

func seed(ctx context.Context, c *redis.Client, scenario string, keys []string, value string, hashArgs []any) error {
	var write func(p redis.Pipeliner, key string)
	switch scenario {
	case "get":
		write = func(p redis.Pipeliner, key string) { p.Set(ctx, key, value, 0) }
	case "hget", "hgetall", "hset":
		write = func(p redis.Pipeliner, key string) { p.HSet(ctx, key, hashArgs...) }
	default:
		return nil
	}
	for i := 0; i < len(keys); i += 500 {
		p := c.Pipeline()
		for _, k := range keys[i:min(i+500, len(keys))] {
			write(p, k)
		}
		if _, err := p.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

type result struct {
	latencies []int64 // ns, one per completed op
	errors    int64
	elapsed   time.Duration
}

// drive runs op from n goroutines until d elapses and records each op's latency.
func drive(n int, d time.Duration, op func(*rand.Rand) error) result {
	var (
		mu    sync.Mutex
		all   []int64
		errs  atomic.Int64
		wg    sync.WaitGroup
		start = time.Now()
		end   = start.Add(d)
	)
	for w := 0; w < n; w++ {
		wg.Add(1)
		go func(seed int64) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(seed))
			local := make([]int64, 0, 1<<14)
			for time.Now().Before(end) {
				t := time.Now()
				if err := op(rng); err != nil {
					errs.Add(1)
					continue
				}
				local = append(local, int64(time.Since(t)))
			}
			mu.Lock()
			all = append(all, local...)
			mu.Unlock()
		}(int64(w) + 1)
	}
	wg.Wait()
	return result{latencies: all, errors: errs.Load(), elapsed: time.Since(start)}
}

func parseSettings(p map[string]any) (settings, error) {
	s := settings{concurrency: 64, keySpace: 1000, payloadBytes: 64, hashFields: 10}
	for k, v := range p {
		switch k {
		case "autopipeline":
			b, ok := v.(bool)
			if !ok {
				return s, fmt.Errorf("autopipeline must be a bool, got %v", v)
			}
			s.autopipeline = b
		case "concurrency":
			if err := posInt(k, v, &s.concurrency); err != nil {
				return s, err
			}
		case "key_space":
			if err := posInt(k, v, &s.keySpace); err != nil {
				return s, err
			}
		case "payload_bytes":
			if err := posInt(k, v, &s.payloadBytes); err != nil {
				return s, err
			}
		case "hash_fields":
			if err := posInt(k, v, &s.hashFields); err != nil {
				return s, err
			}
		}
	}
	return s, nil
}

func posInt(name string, v any, dst *int) error {
	f, ok := v.(float64)
	if !ok || f < 1 || f != float64(int(f)) {
		return fmt.Errorf("%s must be a positive integer, got %v", name, v)
	}
	*dst = int(f)
	return nil
}
