package models

import (
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Comment struct {
	Post_id    int       `json:"post_id"`
	User_id    int       `json:"user_id"`
	Comment_id int       `json:"comment_id"`
	Comment    string    `json:"comment"`
	Created_at time.Time `json:"created_at"`
}
