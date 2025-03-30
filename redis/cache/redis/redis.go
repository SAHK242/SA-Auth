package rediscache

import (
	"context"
	"fmt"

	store "github.com/eko/gocache/store/redis/v4"
	"github.com/redis/go-redis/v9"
)

// NewRedisStore creates a new Redis store
// This should be called only once in the application
func NewRedisStore(client *redis.Client) *store.RedisStore {
	return store.NewRedis(client)
}

// NewRedisClient creates a new Redis client and try to establish a connection
// This should be called only once in the application
func NewRedisClient(opts ...Option) *redis.Client {
	o := eval(opts)

	redisClient := redis.NewClient(&redis.Options{
		Addr:         o.addr,
		MinIdleConns: o.minIdleConns,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Errorf("failed to connect to redis: %w", err))
	}

	return redisClient
}
