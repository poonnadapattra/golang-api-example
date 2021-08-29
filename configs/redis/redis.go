package redis

import (
	"context"

	"example.com/api-example/configs"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     configs.REDIS_ADDRESS,
		Password: configs.REDIS_PASSWORD,
	})

	return rdb
}
