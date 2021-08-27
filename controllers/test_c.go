package controllers

import (
	"log"
	"net/http"
	"time"

	"example.com/api-example/models"
	"github.com/gin-gonic/gin"
)

func (ctrls *Controllers) GetTest(c *gin.Context) {
	var collections []models.Collections
	ctrls.Database.Find(&collections)
	c.JSON(http.StatusOK, gin.H{"message": "OK", "results": collections})
}

func (ctrls *Controllers) GetTestRedis(c *gin.Context) {
	key := c.Param("key")
	val, err := ctrls.RedisService.GetValue(key)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found", "results": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "results": val})
	}
}

func (ctrls *Controllers) SetTestRedis(c *gin.Context) {
	var data = struct {
		Key   string
		Value string
		Time  time.Duration
	}{}
	c.ShouldBind(&data)

	err := ctrls.RedisService.SetValue(data.Key, data.Value, data.Time*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Success"})
}

func (ctrls *Controllers) DeleteTestRedis(c *gin.Context) {
	key := c.Param("key")

	err := ctrls.RedisService.DeleteValue(key)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Success"})
}
