package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create_category(category *models.Category) *models.Category {
	DB.NewRecord(category)
	DB.Create(&category)

	return category
}

func Get_By_Category(category string) []models.Category {
	var categories []models.Category
	DB.Where("category_name=?", category).Find(&categories)
	return categories
}
