package middleware

import (
	"github.com/HazeyamaLab/questack-api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc{
	return func(ctx *gin.Context){
		id, err := util.Parse(ctx)
		if err != nil{
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			ctx.Abort()
			return
		}
		ctx.Set("id", id)
		ctx.Next()
	}
}
