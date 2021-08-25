package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

var (
	PORT           string = ""
	DB_NAME        string = ""
	DB_HOST        string = ""
	DB_USER        string = ""
	DB_PORT        string = ""
	DB_PASSWORD    string = ""
	REDIS_ADDRESS  string = ""
	REDIS_PASSWORD string = ""
)

func InitConstantVariable() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	PORT = geVeriable("PORT")
	DB_HOST = geVeriable("DB_HOST")
	DB_NAME = geVeriable("DB_NAME")
	DB_USER = geVeriable("DB_USER")
	DB_PORT = geVeriable("DB_PORT")
	DB_PASSWORD = geVeriable("DB_PASSWORD")
	REDIS_ADDRESS = geVeriable("REDIS_ADDRESS")
	REDIS_PASSWORD = geVeriable("REDIS_PASSWORD")

}

func geVeriable(key string) string {
	if value, _ := viper.Get(key).(string); value != "" {
		return value
	}
	return os.Getenv(key)
}
