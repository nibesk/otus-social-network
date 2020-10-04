package utils

import (
	"github.com/dgrijalva/jwt-go"
	"service-users/app/config"
)

// Token is JWT claims struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}

func CreateJWTToken(userId uint) string {
	tk := &Token{UserId: userId}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.Env.Server.SessionKey))

	return tokenString
}
