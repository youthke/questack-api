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
	FindMine(ctx *gin.Context)
	FindOneByURL(ctx *gin.Context)
	Update(ctx *gin.Context)
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
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func(s *stackController)FindMine(ctx *gin.Context){
	id := uint(ctx.GetFloat64("id"))

	stacks, err := s.StackService.FindALLByOwnerID(id)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"stacks": stacks,
		"message": "success",
	})
}

func(s *stackController)FindOneByURL(ctx *gin.Context){
	id := ctx.Param("token")
	stack, err := s.StackService.FindOneByID(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"stack": stack,
		"message": "success",
	})
}


func(s *stackController)Update(ctx *gin.Context){
	ownerID:= uint(ctx.GetFloat64("id"))

	var form presenter.StackUpdateForm

	err := ctx.BindJSON(&form)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = s.StackService.Update(form.ID,form.Name,ownerID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
