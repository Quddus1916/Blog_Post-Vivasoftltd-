package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Category struct {
	Category_id   int       `json:"category_id" gorm:"primary_key;auto_increment;not_null"`
	Post_id       int       `json:"post_id"`
	Category_name string    `json:"category_name"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
