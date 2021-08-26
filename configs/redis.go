package configs

import (
	"github.com/go-redis/redis/v8"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     REDIS_ADDRESS,
		Password: REDIS_PASSWORD,
	})

	return rdb
}
