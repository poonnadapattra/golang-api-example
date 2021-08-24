package main

import (
	"log"
	"net/http"
	"os"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = configs.PORT // Default port if not specified
	}

	log.Println("Server Running on Port: ", port)
	http.ListenAndServe(":"+port, router)
}
