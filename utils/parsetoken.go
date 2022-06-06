package utils

import (
	"blogpost.com/config"
	"blogpost.com/types"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	//"time"
)

func ParserToken(usertoken string) string {
	//var claims *types.SignedDetails
	var claims = &types.SignedDetails{}
	token, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Getconfig().SecretKey), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println(token)
	}
	// do something with decoded claims
	return claims.Email
}
