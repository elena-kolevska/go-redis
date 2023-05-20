package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	options := &redis.Options{
		Addr: "localhost:6379",
	}
	rdb := redis.NewClient(options)

	ctx := context.Background()

	args1 := rdb.B().BfAdd().Key("key").Item("element").GetArgs()
	r1, err := rdb.Do(ctx, args1...).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(r1)

	args2 := rdb.B().BfExists().Key("key").Item("element").GetArgs()
	r2, err := rdb.Do(ctx, args2...).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(r2)
}
