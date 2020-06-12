package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const key = "dev"

func GenerateOwnerJwt(ownerID uint) (string, error){

	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"id": ownerID,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}
	return token.SignedString([]byte(key))
}

func Parse(ctx *gin.Context) (float64, error){

	token, err := request.ParseFromRequest(ctx.Request, request.OAuth2Extractor,
		func(token *jwt.Token) (i interface{}, e error) {
			b := []byte(key)
			return b, nil
		})
	if err != nil {
		log.Println(err)
		return 0, err
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	return claims["id"].(float64), nil

}
