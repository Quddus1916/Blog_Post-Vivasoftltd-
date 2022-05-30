package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Blog struct {
	Userid     int       `json:"userid"`
	Postid     int       `json:"postid"`
	Post       string    `json:"post"`
	Category   string    `json:"category"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
