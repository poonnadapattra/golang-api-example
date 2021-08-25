package controllers

import (
	"log"
	"net/http"

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
	val, _ := ctrls.RedisService.GetValue(key)
	c.JSON(http.StatusOK, gin.H{"message": "OK", "results": val})
}

func (ctrls *Controllers) SetTestRedis(c *gin.Context) {
	var data = map[string]string{}
	err := c.ShouldBind(&data)
	err = ctrls.RedisService.SetValue(data["key"], data["value"])
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "SUCCESS"})
}
