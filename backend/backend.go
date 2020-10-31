package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Hello World")

	go connectToRedis()

	for {

	}
	wg.Wait()
}

var ctx = context.Background()

func connectToRedis() {
	db := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "temp_pass_123", // no password set
		DB:       0,               // use default DB
	})

	err := db.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := db.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := db.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
