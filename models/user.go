package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `Json:"password"`
	Token         string    `json:"token"`
	Refresh_token string    `json:"refresh_token"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
