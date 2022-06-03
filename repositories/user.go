package repositories

import (
	"blogpost.com/models"
	"blogpost.com/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func CreateUser(user *models.User) *models.User {
	DB.NewRecord(user)
	DB.Create(&user)

	return user
}

func GetByEmail(email string) *models.User {
	var sigleuser = new(models.User)
	DB.Where("email=?", email).Find(sigleuser)
	return sigleuser

}

func SetToken(user *models.User) bool {

	db := DB.Model(&user).Updates(map[string]interface{}{"token": user.Token, "refresh_token": user.Refresh_token})
	if db == nil {
		return false
	}
	return true

}

func UpdateUser(user *models.User) *models.User {
	user.Password = utils.Encrypt(user.Password)
	DB.Model(&user).Where("id = ?", user.Id).Updates(map[string]interface{}{"name": user.Name, "updated_at": user.Updated_at, "password": user.Password})
	return user
}

func LogOut(user *models.User) error {
	err := DB.Model(&user).Where("id = ?", user.Id).Updates(map[string]interface{}{"token": " ", "refresh_token": " ", "updated_at": user.Updated_at}).Error
	return err
}

func UploadPic(user *models.User) error {
	err := DB.Model(&user).Where("id = ?", user.Id).Update("image_path", user.Image_path).Error
	return err
}

func GetUserDetails(id int) *models.User {
	var sigleuser = new(models.User)
	DB.Where("id=?", id).Find(sigleuser)
	return sigleuser

}
