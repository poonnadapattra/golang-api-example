package routers

import (
	"example.com/api-example/controllers"
	"github.com/gin-gonic/gin"
)

func SetCollectionRoutes(router *gin.RouterGroup, ctrls *controllers.Controllers) {

	router.GET("collections", ctrls.GetCollection)
	router.GET("collections/:id", ctrls.GetCollectionById)
	router.POST("collections", ctrls.CreateCollection)
	router.PUT("collections/:id", ctrls.PutCollection)
	router.PATCH("collections", ctrls.UpdateCollection)
	router.DELETE("collections/:id", ctrls.DeleteCollection)

}
