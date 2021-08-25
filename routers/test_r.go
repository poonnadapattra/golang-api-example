package routers

import (
	"example.com/api-example/controllers"
	"github.com/gin-gonic/gin"
)

func SetTestRoutes(router *gin.RouterGroup, ctrls *controllers.Controllers) {

	router.GET("test", ctrls.GetTest)
	router.GET("test/redis/:key", ctrls.GetTestRedis)
	router.POST("test/redis", ctrls.SetTestRedis)
}
