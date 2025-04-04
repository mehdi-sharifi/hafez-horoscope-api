package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"hafez-horoscope-api/config"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func GetRedisClient(config *config.Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		DB:   config.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{
		client: client,
	}, nil
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return rc.client.Get(ctx, key).Result()
}

func (rc *RedisClient) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return rc.client.Set(ctx, key, value, expiration).Err()
}
