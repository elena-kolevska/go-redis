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

	cmd1 := rdb.B().BfAdd().Key("key").Item("element").Build()
	r1, err := rdb.Run(ctx, cmd1.GetArgs()).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(r1)

	cmd2 := rdb.B().BfExists().Key("key").Item("element").Build()
	r2, err := rdb.Run(ctx, cmd2.GetArgs()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(r2)
}
