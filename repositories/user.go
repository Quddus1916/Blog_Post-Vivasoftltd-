package repositories

import (
	"blogpost.com/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create_user(user *models.User) *models.User {
	DB.NewRecord(user)
	DB.Create(&user)

	return user
}

func Get_by_email(email string) *models.User {
	var sigleuser = new(models.User)
	DB.Where("email=?", email).Find(sigleuser)
	return sigleuser

}

func Set_token(user *models.User) bool {

	db := DB.Model(&user).Updates(map[string]interface{}{"token": user.Token, "refreshtoken": user.Refreshtoken})
	if db == nil {
		return false
	}
	return true

}
