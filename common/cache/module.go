package cache

import (
	"context"
	"fmt"

	"auth/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"

	client "auth/redis/cache/redis"
)

type props struct {
	fx.In
	Config config.Config
	Logger *zap.SugaredLogger
	Redis  *redis.Client
}

func NewRedis(lifecycle fx.Lifecycle, config config.Config) *redis.Client {
	host := config.MustGetStringSlice("REDIS_HOST")
	port := config.MustGetInt("REDIS_PORT")
	addr := fmt.Sprintf("%s:%d", host, port)

	redisClient := client.NewRedisClient(
		client.WithAddr(addr),
		client.WithMinIdleConns(config.GetInt("REDIS_MIN_CONN", 2)),
	)

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return redisClient.Close()
		},
	})

	return redisClient
}

var Module = fx.Provide(
	NewRedis,
)
