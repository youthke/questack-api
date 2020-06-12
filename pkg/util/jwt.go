package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateOwnerJwt(ownerID uint) (string, error){

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"id": ownerID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	return token.SignedString([]byte("dev"))
}

