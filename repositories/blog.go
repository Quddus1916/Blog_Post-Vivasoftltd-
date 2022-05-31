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

func Get_all() []models.Blog {
	var blogs []models.Blog
	DB.Find(&blogs)
	return blogs
}

func Update_blog(blog *models.Blog) *models.Blog {
	DB.Model(&blog).Where("post_id = ?", blog.Post_id).Updates(map[string]interface{}{"post": blog.Post, "updated_at": blog.Updated_at})
	return blog
}

func Delete_blog(blog *models.Blog) string {
	DB.Where("post_id = ?", blog.Post_id).Delete(&blog)
	return "delete successful"
}

func Post_comment(comment *models.Comment) *models.Comment {
	DB.NewRecord(comment)
	DB.Create(&comment)
	return comment
}
