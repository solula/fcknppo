package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"waterfall-backend/internal/modules/features/config"
)

func NewClient(cfg config.Config) (*redis.Client, error) {
	c := redis.NewClient(&redis.Options{
		Addr:         getAddress(cfg),
		Password:     "",
		DB:           0,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		PoolSize:     1000,
	})

	_, err := c.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к redis: %w", err)
	}

	return c, nil
}

func getAddress(cfg config.Config) string {
	return fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
}
