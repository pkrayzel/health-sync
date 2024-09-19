package api

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

// InitRedis initializes the Redis connection
func InitRedis() {
	redisURL := os.Getenv("REDIS_URL") // Ensure REDIS_URL does not contain redis://

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: os.Getenv("REDIS_PASSWORD"), // Redis password from the environment
		DB:       0,                           // Use default DB
	})

	// Test the connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	} else {
		log.Printf("Connected to Redis at %s: %s", redisURL, pong)
	}
}
