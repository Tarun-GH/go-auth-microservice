package config

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient(redisHost string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: redisHost, //"localhost:6379"
	})
}
