package main

import (
	"log"
	"net/http"

	"example.com/api-example/configs"
	"example.com/api-example/routers"
	dbService "example.com/api-example/services/db"
	redisService "example.com/api-example/services/redis"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	configs.InitConstantVariable()
	redis := redisService.InitRedis()
	db, _ := dbService.InitDatabase()
	router = gin.Default()
	routers.InitRouters(router, db, redis)
}

func main() {
	log.Println("Server Running on Port: ", configs.PORT)
	http.ListenAndServe(":"+configs.PORT, router)
}
