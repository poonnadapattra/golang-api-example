package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouters(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	SetCollectionRoutes(api, db)
}
