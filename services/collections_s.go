package services

import (
	"time"

	"example.com/api-example/models"
	"gorm.io/gorm"
)

func GetCollection(db *gorm.DB, query map[string]interface{}) <-chan []models.Collections {
	c := make(chan []models.Collections)

	go func() {
		defer close(c)

		time.Sleep(time.Second * 3)
		var collections []models.Collections
		db.Where(query).Find(&collections)
		c <- collections
	}()

	return c
}
