package core

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

// NewRedis creates a new connector to Redis
func NewRedis(host, port, password string) (*redis.Client, error) {
	// Forming a Redis address
	addr := fmt.Sprintf("%s:%s", host, port)

	// Creating a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // Blank password means no password
		DB:       0,        // Default is DB 0
	})

	// Checking the connection
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("не удалось подключиться к Redis: %v", err)
	}

	return client, nil
}
