package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Blog struct {
	User_id int    `json:"user_id"`
	Post_id int    `json:"post_id"`
	Post    string `json:"post"`
	//Category   string    `json:"category"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
