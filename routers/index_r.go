package routers

import (
	"net/http"

	"example.com/api-example/controllers"
	authService "example.com/api-example/services/auth"
	redisServices "example.com/api-example/services/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func InitRouters(r *gin.Engine, db *gorm.DB, redis *redis.Client) {

	jwtService := authService.JWTAuthService()
	redisService := redisServices.Redis{Redis: redis}
	authAtrls := &controllers.Auth{
		Database:    db,
		AuthService: jwtService,
	}
	ctrls := &controllers.Controllers{
		Database:     db,
		RedisService: redisService,
	}

	api := r.Group("/api")

	SetAuth(r, authAtrls)
	api.Use(authAtrls.AuthorizeJWT())

	SetCollectionRoutes(api, ctrls)
	SetTestRoutes(api, ctrls)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}
