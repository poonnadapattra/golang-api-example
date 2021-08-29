package routers

import (
	"example.com/api-example/controllers"
	"github.com/gin-gonic/gin"
)

func SetAuth(router *gin.Engine, ctrls *controllers.Auth) {

	router.POST("/login", ctrls.Login)
}
