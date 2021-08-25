package services

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Redis struct {
	Redis *redis.Client
}

type RedisDataStruct struct {
	Key   string
	Value string
}

func (redis *Redis) GetValue(key string) (val string, err error) {
	val, err = redis.Redis.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (redis *Redis) SetValue(key string, value string) (err error) {
	err = redis.Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
