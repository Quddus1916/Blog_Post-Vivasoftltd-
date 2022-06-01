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

	db := DB.Model(&user).Updates(map[string]interface{}{"token": user.Token, "refresh_token": user.Refresh_token})
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

func Log_out(user *models.User) error {
	err := DB.Model(&user).Where("id = ?", user.Id).Updates(map[string]interface{}{"token": " ", "refresh_token": " ", "updated_at": user.Updated_at}).Error
	return err
}

func Upload_pro_pic(user *models.User) error {
	err := DB.Model(&user).Where("id = ?", user.Id).Update("image_path", user.Image_path).Error
	return err
}

func Get_userdetails(id int) *models.User {
	var sigleuser = new(models.User)
	DB.Where("id=?", id).Find(sigleuser)
	return sigleuser

}
