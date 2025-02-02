package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"weather_api/internal/model"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client  *redis.Client
	expires time.Duration
}

func NewRedisCache(url string, expires int) (*RedisCache, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)
	return &RedisCache{
		client:  client,
		expires: time.Duration(expires) * time.Second,
	}, nil
}

func (c *RedisCache) Set(ctx context.Context, key string, value *model.WeatherData) error {
	json, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, json, c.expires).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string) (*model.WeatherData, error) {
	val, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var weather model.WeatherData
	if err := json.Unmarshal([]byte(val), &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
