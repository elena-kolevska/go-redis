package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	options := &redis.Options{
		Addr: "localhost:6379",
	}
	rdb := redis.NewClient(options)

	ctx := context.Background()

	type User struct {
		Name   string   `json:"name"`
		Age    int      `json:"age"`
		Titles []string `json:"titles"`
	}

	// Save a user in Redis
	userObj := &User{
		Name:   "John Doe",
		Age:    30,
		Titles: []string{"Mr", "Sir"},
	}

	userJson, err := json.Marshal(userObj)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := rdb.B().JsonSet().Key("json:user:1").Path("$").Value(userJson).GetArgs()
	fmt.Println(rdb.Do(ctx, args...))

	// Get a user from Redis
	args = rdb.B().JsonGet().Key("json:user:1").Path("$").GetArgs()
	userFromDbJSON, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	var usersFromDb []User
	err = json.Unmarshal([]byte(userFromDbJSON.(string)), &usersFromDb)
	fmt.Println(usersFromDb[0])
}
