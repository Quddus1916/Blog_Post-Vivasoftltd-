package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create_Blog(blog *models.Blog) *models.Blog {
	DB.NewRecord(blog)
	DB.Create(&blog)

	return blog
}

func Get_By_Category(category string) []models.Blog {
	var blogs []models.Blog
	DB.Where("category=?", category).Find(&blogs)
	return blogs
}
