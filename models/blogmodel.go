package models

import (
	"blogpost.com/database"
	"github.com/jinzhu/gorm"
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

var db2 *gorm.DB

func init() {
	database.Connect()

	db2 = database.GetDB()
	//table name is defined through this
	db2.AutoMigrate(&Blog{})
}

func CreateBlog(blog *Blog) *Blog {
	db.NewRecord(blog)
	db.Create(&blog)

	return blog
}

func Getblogbycategory(category string) []Blog {
	var blogs []Blog
	db.Where("category=?", category).Find(&blogs)
	return blogs
}
