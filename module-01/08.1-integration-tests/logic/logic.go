package logic

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func SetValue(ctx context.Context, c Client, val string) string {
	err := c.Set(ctx, "A", val, time.Hour).Err()
	if err != nil {
		panic(err)
	}
	res := c.Get(ctx, "A")
	if res.Err() != nil {
		panic(err)
	}
	return res.Val()
}

func GetClient(dsn string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: dsn,
	})
}
