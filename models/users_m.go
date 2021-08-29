package models

import (
	"time"
)

type Users struct {
	Id          int64     `json:"id" orm:"auto" gorm:"primary_key"`
	Username    string    `json:"username" orm:"size(128)"`
	Password    string    `json:"password" orm:"size(64)"`
	CreatedDate time.Time `json:"created_date"`
}
