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

	// Create a few documents.
	rdb.HSet(ctx, "users:1", "full_name", "John Doe", "age", 30)
	rdb.HSet(ctx, "users:2", "full_name", "Jane Doe", "age", 33)
	rdb.HSet(ctx, "users:3", "full_name", "John Malkovich", "age", 56)

	argsCreate := rdb.B().FtCreate().
		Index("myindex").
		OnHash().
		Prefix(1).Prefix("users:").
		Schema().
		FieldName("full_name").As("name").Text().
		FieldName("age").As("age").Numeric().
		GetArgs()

	fmt.Println(rdb.Do(ctx, argsCreate...))

	argsSearch := rdb.B().FtSearch().
		Index("myindex").
		Query("@name:John @age: [25, 65]").
		GetArgs()
	fmt.Println(rdb.Do(ctx, argsSearch...))
}
