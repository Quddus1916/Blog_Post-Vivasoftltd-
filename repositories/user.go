package repositories

import (
	"blogpost.com/models"
	"blogpost.com/utils"
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

func Update_user(user *models.User) *models.User {
	user.Password = utils.Encrypt(user.Password)
	DB.Model(&user).Where("id = ?", user.Id).Updates(map[string]interface{}{"name": user.Name, "updated_at": user.Updated_at, "password": user.Password})
	return user
}
