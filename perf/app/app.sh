#!/usr/bin/env bash
set -euo pipefail

case "${1:-}" in
  list-scenarios)
    printf '%s\n' set get hset hget hgetall
    ;;
  setup)
    repo=$(git rev-parse --show-toplevel 2>/dev/null || true)
    cp -R . "$BENCH_WORKDIR/src"
    cd "$BENCH_WORKDIR/src"
    [ -f go.mod ] || go mod init goredis-bench

    v=$BENCH_LIBRARY_VERSION
    lib=""
    case "$v" in
      path:*) lib=${v#path:} ;;
      *)
        # A tag or SHA of the checkout this app lives in is exported into the
        # workdir, so a commit that only exists in a fork needs no proxy lookup.
        if [ -n "$repo" ] \
          && sha=$(git -C "$repo" rev-parse --verify --quiet "$v^{commit}") \
          && [ "$(git -C "$repo" show "$sha:go.mod" 2>/dev/null | sed -n 1p)" = "module github.com/redis/go-redis/v9" ]; then
          lib="$BENCH_WORKDIR/lib"
          mkdir -p "$lib"
          git -C "$repo" archive "$sha" | tar -x -C "$lib"
        fi
        ;;
    esac

    if [ -n "$lib" ]; then
      go mod edit -replace "github.com/redis/go-redis/v9=$lib"
    else
      go get "github.com/redis/go-redis/v9@$v"
    fi
    go mod tidy
    go build -o "$BENCH_WORKDIR/goredis-app" .
    ;;
  run)
    exec "$BENCH_WORKDIR/goredis-app"
    ;;
  *)
    echo "goredis: unknown verb '${1:-}'" >&2
    exit 64
    ;;
esac
