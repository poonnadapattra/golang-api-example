package controllers

import (
	authService "example.com/api-example/services/auth"
	redisService "example.com/api-example/services/redis"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

type RedisController struct {
	Redis *redis.Client
}

type Controllers struct {
	Database     *gorm.DB
	RedisService redisService.Redis
}

type Auth struct {
	AuthService authService.Auth
	Database    *gorm.DB
}
