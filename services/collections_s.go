package services

import (
	"example.com/api-example/models"
	"gorm.io/gorm"
)

func GetCollection(db *gorm.DB, query map[string]interface{}) <-chan []models.Collections {
	c := make(chan []models.Collections)

	go func() {
		defer close(c)

		var collections []models.Collections
		db.Where(query).Find(&collections)

		for i, _ := range collections {
			db.Model(collections[i]).Association("Groups").Find(&collections[i].Groups)
		}

		c <- collections
	}()

	return c
}
