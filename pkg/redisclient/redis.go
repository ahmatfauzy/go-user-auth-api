package redisclient

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{client: rdb}
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisClient) Del(key string) error {
	return r.client.Del(ctx, key).Err()
}