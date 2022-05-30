package models

import (
	//"blogpost.com/database"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `Json:"password"`
	Token        string `json:"token"`
	Refreshtoken string `json:"refreshtoken"`
}
