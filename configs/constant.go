package configs

import (
	"os"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type DBController struct {
	Database *gorm.DB
}

var (
	PORT           string = "9999"
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

	}

	PORT = geVeriable("PORT", PORT)
	DB_HOST = geVeriable("DB_HOST", DB_HOST)
	DB_NAME = geVeriable("DB_NAME", DB_HOST)
	DB_USER = geVeriable("DB_USER", DB_USER)
	DB_PORT = geVeriable("DB_PORT", DB_PORT)
	DB_PASSWORD = geVeriable("DB_PASSWORD", DB_PASSWORD)
	REDIS_ADDRESS = geVeriable("REDIS_ADDRESS", REDIS_ADDRESS)
	REDIS_PASSWORD = geVeriable("REDIS_PASSWORD", REDIS_PASSWORD)

}

func geVeriable(key string, val string) string {
	if value, _ := viper.Get(key).(string); value != "" {
		return value
	} else if value = os.Getenv(key); value != "" {
		return value
	}
	return val
}
