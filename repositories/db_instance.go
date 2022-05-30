package repositories

import (
	"blogpost.com/database"
	"blogpost.com/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	DB = database.GetDB()
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Blog{})
}
