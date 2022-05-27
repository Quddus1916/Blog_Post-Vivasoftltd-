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
