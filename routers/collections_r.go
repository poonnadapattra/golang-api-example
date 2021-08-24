package routers

import (
	"example.com/api-example/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetCollectionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ctrls := controllers.DBController{Database: db}

	router.GET("collections", ctrls.GetCollection)
	router.GET("collections/:id", ctrls.GetCollectionById)
	router.POST("collections", ctrls.CreateCollection)
	router.PATCH("collections", ctrls.UpdateCollection)
	router.DELETE("collections/:id", ctrls.DeleteCollection)
}
