package helpers

import (
	"blogpost.com/config"
	"blogpost.com/types"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Verify_token(usertoken string) (bool, error) {
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
