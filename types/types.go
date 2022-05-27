package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email    string
	Username string
	User_id  int
	jwt.StandardClaims
}

type UserLogIn struct {
	Email    string
	Password string
}

type Category struct {
	Category string
}
