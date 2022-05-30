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

func Get_all() []models.Blog {
	var blogs []models.Blog
	DB.Find(&blogs)
	return blogs
}

func Update_blog(blog *models.Blog) *models.Blog {
	DB.Model(&blog).Where("postid = ?", blog.Postid).Updates(map[string]interface{}{"post": blog.Post, "updated_at": blog.Updated_at, "category": blog.Category})
	return blog
}

func Delete_blog(blog *models.Blog) string {
	DB.Where("postid = ?", blog.Postid).Delete(&blog)
	return "delete successful"
}
