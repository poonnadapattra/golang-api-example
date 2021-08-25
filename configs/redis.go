package configs

import (
	"github.com/go-redis/redis/v8"
)

// var ctx = context.Background()

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Network:  "",
		Addr:     REDIS_ADDRESS,
		Password: REDIS_PASSWORD,
	})

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	log.Fatalln('1', err)
	// }

	// val, err := rdb.Get(ctx, "mykey").Result()
	// if err != nil {
	// 	log.Fatalln('2', err)
	// }
	// fmt.Println("key", val)

	return rdb
}
