package token

import (
	"blogpost.com/config"
	"blogpost.com/types"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func Generate_tokens(email string, username string, password string) (signed_token string, signed_refreshtoken string, err error) {

	claims := &types.SignedDetails{
		Email:         email,
		User_name:     username,
		User_Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.Getconfig().SecretKey))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.Getconfig().SecretKey))

	return token, refresh_token, nil
}
