package controller

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/service"
	"github.com/HazeyamaLab/questack-api/pkg/presenter"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StackController interface {
	Create(ctx *gin.Context)
}

type stackController struct {
	service.StackService
}

func NewStackController(s service.StackService)StackController{
	return &stackController{s}
}

func(s *stackController)Create(ctx *gin.Context){
	id := uint(ctx.GetFloat64("id"))


	var form presenter.StackCreateForm
	err := ctx.BindJSON(&form)

	if err != nil{
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = s.StackService.Create(form.Name, uint(id))

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
