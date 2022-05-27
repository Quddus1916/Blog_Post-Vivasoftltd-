package helpers

import (
	"blogpost.com/config"
	"blogpost.com/database"
	"blogpost.com/models"
	"blogpost.com/types"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	database.Connect()

	db = database.GetDB()
	//table name is defined through this
	db.AutoMigrate(&models.User{})
}

func Getuserdetails(id int) *models.User {
	var sigleuser = new(models.User)
	db.Where("id=?", id).Find(sigleuser)
	return sigleuser

}

func Verifytoken(usertoken string) (bool, error) {
	claims := &types.SignedDetails{}
	flag := false

	fmt.Println(config.Getconfig().SecretKey)

	token, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Getconfig().SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return flag, err
		}

		return flag, err
	}
	if !token.Valid {

		return flag, err
	}
	flag = true

	return flag, nil

}
