package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang-app/storage"
	"time"
)

func main() {
	cfg := storage.Config{
		Addr:        "redis:6379",
		User:        "default",
		Password:    "my-password",
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 10 * time.Second,
		Timeout:     5 * time.Second,
	}

	db, err := storage.NewClient(context.Background(), cfg)

	if err != nil {
		panic(err)
	}

	if err := db.Set(context.Background(), "key", "test value", 0).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	//30sec for life in DB
	if err := db.Set(context.Background(), "key2", 111, 30*time.Second).Err(); err != nil {
		fmt.Printf("failed to set data, error: %s", err.Error())
	}

	val, err := db.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Printf("failed to get value, error: %v\n", err)
	}

	val2, err := db.Get(context.Background(), "key2").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Printf("failed to get value, error: %v\n", err)
	}

	fmt.Printf("value: %v\n", val)
	fmt.Printf("value2: %v\n", val2)

	time.Sleep(40 * time.Second)

	oldVal2, err := db.Get(context.Background(), "key2").Result()
	if err == redis.Nil {
		fmt.Println("value not found")
	} else if err != nil {
		fmt.Printf("failed to get value, error: %v\n", err)
	}

	fmt.Printf("after some time value : %v\n", oldVal2)

	fmt.Println("completed.")
}
