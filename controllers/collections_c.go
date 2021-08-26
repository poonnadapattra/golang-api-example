package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/api-example/models"
	"github.com/gin-gonic/gin"
)

func (ctrls *Controllers) GetCollection(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}
	from := "database"

	var collections []models.Collections
	val, err := ctrls.RedisService.GetValue("get_collections_x")

	if err != nil {

		if _type != "" {
			_where["type"] = _type
		}

		ctrls.Database.Where(_where).Find(&collections)

		for i, _ := range collections {
			ctrls.Database.Model(collections[i]).Association("Groups").Find(&collections[i].Groups)
		}

		// data, _ := json.Marshal(collections)
		// ctrls.RedisService.SetValue("get_collections", string(data), 5*time.Minute)

	} else {
		from = "redis"
		json.Unmarshal([]byte(val), &collections)
	}

	c.JSON(http.StatusOK, gin.H{"form": from, "count": len(collections), "results": &collections})

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
	ctrls.Database.Delete(&collections, id)

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
}
