package redis_services

import (
	"context"
	"log"
	"time"

	"example.com/api-example/configs"
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

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     configs.REDIS_ADDRESS,
		Password: configs.REDIS_PASSWORD,
	})

	return rdb
}

func (redis *Redis) GetValue(key string) (val string, err error) {
	val, err = redis.Redis.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return
}

func (redis *Redis) SetValue(key string, value string, s time.Duration) (err error) {
	err = redis.Redis.Set(ctx, key, value, s).Err()
	if err != nil {
		log.Println(err)
	}
	return
}

func (redis *Redis) DeleteValue(key string) (err error) {
	err = redis.Redis.Del(ctx, key).Err()
	if err != nil {
		log.Println(err)
	}
	return
}
