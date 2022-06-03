package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreatePost(blog *models.Post) *models.Post {
	DB.NewRecord(blog)
	DB.Create(&blog)

	return blog
}

func GetAll() []models.Post {
	var blogs []models.Post
	DB.Find(&blogs)
	return blogs
}

func UpdatePost(post *models.Post) *models.Post {
	DB.Model(&post).Where("id = ?", post.Id).Updates(map[string]interface{}{"post": post.Post, "updated_at": post.Updated_at})
	return post
}

func DeletePost(post *models.Post) string {
	DB.Where("id = ?", post.Id).Delete(&post)
	return "delete successful"
}
