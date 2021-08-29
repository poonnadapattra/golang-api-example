package main

import (
	"log"
	"net/http"

	"example.com/api-example/configs"
	database "example.com/api-example/configs/db"
	redis "example.com/api-example/configs/redis"
	"example.com/api-example/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	configs.InitConstantVariable()
	redis := redis.InitRedis()
	db, _ := database.InitDatabase()
	router = gin.Default()
	routers.InitRouters(router, db, redis)
}

func main() {
	log.Println("Server Running on Port: ", configs.PORT)
	http.ListenAndServe(":"+configs.PORT, router)
}
