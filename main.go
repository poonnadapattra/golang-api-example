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
	db, _ := configs.InitDatabase()
	router = gin.Default()
	routers.InitRouters(router, db)
}

func main() {
	log.Println("Server Running on Port: ", 8080)
	http.ListenAndServe(":8080", router)
}
