package redis

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"

	"home-service/pkg/config"
)

type RedisAdapter struct {
	client *redis.Client
	ttl    time.Duration
	cfg    config.RedisConfig
	log    slog.Logger
}

func NewRedisCache(cfg config.RedisConfig, log slog.Logger) (*RedisAdapter, error) {
	op := "adapters.cache.redis"
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		log.Error("%s: %w", op, err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &RedisAdapter{
		client: client,
		ttl:    1 * time.Hour,
		cfg:    cfg,
		log:    log,
	}, nil
}
