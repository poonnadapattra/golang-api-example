package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"example.com/api-example/models"
	s "example.com/api-example/services"
	"github.com/gin-gonic/gin"
)

func (ctrls *Controllers) GetCollection(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}
	from := "database"

	var collections []models.Collections
	val, _ := ctrls.RedisService.GetValue("get_collections")
	whereVal, _ := ctrls.RedisService.GetValue("get_collections_where")
	whereData, _ := json.Marshal(_where)

	if whereVal != "" || string(whereData) != whereVal {

		if _type != "" {
			_where["type"] = _type
		}

		c := s.GetCollection(ctrls.Database, _where)
		collections = <-c

		data, _ := json.Marshal(collections)
		ctrls.RedisService.SetValue("get_collections", string(data), 10*time.Minute)
		ctrls.RedisService.SetValue("get_collections_where", string(whereData), 10*time.Minute)

	} else {
		from = "redis"
		json.Unmarshal([]byte(val), &collections)
	}

	c.JSON(http.StatusOK, gin.H{"from": from, "count": len(collections), "results": &collections})

}

func (ctrls *Controllers) GetCollectionById(c *gin.Context) {
	id := c.Param("id")
	var collections models.Collections

	ctrls.Database.First(&collections, id)
	ctrls.Database.Model(&collections).Association("Groups").Find(&collections.Groups)

	c.JSON(http.StatusOK, gin.H{"results": &collections})
}

func (ctrls *Controllers) CreateCollection(c *gin.Context) {
	var collection models.Collections
	err := c.ShouldBind(&collection)

	result := ctrls.Database.Create(&collection)
	for i, _ := range collection.Groups {
		collection.Groups[i].CollectionId = collection.Id
		ctrls.Database.Create(collection.Groups[i])
	}

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &collection})
	}
}

func (ctrls *Controllers) PutCollection(c *gin.Context) {
	id := c.Param("id")
	var collection models.Collections

	ctrls.Database.First(&collection, id)
	err := c.ShouldBind(&collection)

	result := ctrls.Database.Save(&collection)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &collection})
	}
}

func (ctrls *Controllers) UpdateCollection(c *gin.Context) {

	var collection models.Collections
	err := c.ShouldBind(&collection)

	result := ctrls.Database.Updates(collection)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &collection})
	}
}

func (ctrls *Controllers) DeleteCollection(c *gin.Context) {
	id := c.Param("id")
	var collections models.Collections
	result := ctrls.Database.Delete(&collections, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
