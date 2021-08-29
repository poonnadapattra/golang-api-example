package database

import (
	"fmt"

	"example.com/api-example/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (db *gorm.DB, err error) {

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", configs.DB_HOST, configs.DB_USER, configs.DB_NAME, configs.DB_PORT, configs.DB_PASSWORD)

	db, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	return
}
