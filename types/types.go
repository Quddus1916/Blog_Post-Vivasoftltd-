package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email         string
	User_name     string
	User_Password string
	jwt.StandardClaims
}

type UserLogIn struct {
	Email    string
	Password string
}

type Category struct {
	Category string
}

type UserId struct {
	Id int
}
