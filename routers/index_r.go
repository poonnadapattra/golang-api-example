package routers

import (
	"net/http"

	"example.com/api-example/controllers"
	"example.com/api-example/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func InitRouters(r *gin.Engine, db *gorm.DB, redis *redis.Client) {

	redisService := services.Redis{Redis: redis}
	ctrls := &controllers.Controllers{
		Database:     db,
		RedisService: redisService,
	}

	api := r.Group("/api")
	SetCollectionRoutes(api, db)
	SetTestRoutes(api, ctrls)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})
}
