package controller

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/service"
	"github.com/HazeyamaLab/questack-api/pkg/presenter"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QuestionController interface {
	Create(ctx *gin.Context)
}

type questionController struct {
	service.QuestionService
}

func NewQuestionController(s service.QuestionService) QuestionController{
	return &questionController{s}
}

func(q *questionController)Create(ctx *gin.Context){
	var form presenter.QuestionCreateForm

	err := ctx.BindJSON(&form)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	err = q.QuestionService.Create(form.StackURL, form.Title, form.Author, form.Content)

	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})


}
