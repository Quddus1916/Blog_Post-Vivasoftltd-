package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func PostComment(comment *models.Comment) *models.Comment {
	DB.NewRecord(comment)
	DB.Create(&comment)
	return comment
}

func UpdateComment(comment *models.Comment) *models.Comment {
	DB.Model(&comment).Where("id = ?", comment.Id).Updates(map[string]interface{}{"comment": comment.Comment})
	return comment
}

func DeleteComment(comment *models.Comment) string {
	DB.Where("id = ?", comment.Id).Delete(&comment)
	return "delete successful"
}

func GetCommentDetails(id int) *models.Comment {
	var comment = new(models.Comment)
	DB.Where("id=?", id).Find(comment)
	return comment

}
