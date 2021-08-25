package main

import (
	"log"
	"net/http"

	configs "example.com/api-example/configs"
	routers "example.com/api-example/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	configs.InitConstantVariable()
	redis := configs.InitRedis()
	db, _ := configs.InitDatabase()
	router = gin.Default()
	routers.InitRouters(router, db, redis)
}

func main() {
	log.Println("Server Running on Port: ", configs.PORT)
	http.ListenAndServe(":"+configs.PORT, router)
}
