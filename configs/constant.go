package configs

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

var (
	DB_NAME     string = ""
	DB_HOST     string = ""
	DB_USER     string = ""
	DB_PORT     string = ""
	DB_PASSWORD string = ""
)

func InitConstantVariable() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	DB_HOST, _ = viper.Get("DB_HOST").(string)
	DB_NAME, _ = viper.Get("DB_NAME").(string)
	DB_USER, _ = viper.Get("DB_USER").(string)
	DB_PORT, _ = viper.Get("DB_PORT").(string)
	DB_PASSWORD, _ = viper.Get("DB_PASSWORD").(string)
}
