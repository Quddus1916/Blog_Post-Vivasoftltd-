package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Post_comment(comment *models.Comment) *models.Comment {
	DB.NewRecord(comment)
	DB.Create(&comment)
	return comment
}

func Update_comment(comment *models.Comment) *models.Comment {
	DB.Model(&comment).Where("comment_id = ?", comment.Comment_id).Updates(map[string]interface{}{"comment": comment.Comment})
	return comment
}

func Delete_comment(comment *models.Comment) string {
	DB.Where("post_id = ?", comment.Post_id).Delete(&comment)
	return "delete successful"
}

func Get_comment_details(id int) *models.Comment {
	var comment = new(models.Comment)
	DB.Where("comment_id=?", id).Find(comment)
	return comment

}
