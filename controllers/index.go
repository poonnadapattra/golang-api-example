package controllers

import (
	"example.com/api-example/services"
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
	RedisService services.Redis
}
