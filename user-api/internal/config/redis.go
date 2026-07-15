package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")     // e.g. "localhost:6379"
	pass := os.Getenv("REDIS_PASSWORD") // optional
	db := 0                             // default DB

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})

	// Test connection
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connection established")
	return rdb
}
