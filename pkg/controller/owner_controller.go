package controller

import (
	"github.com/HazeyamaLab/questack-api/pkg/domain/service"
	"github.com/HazeyamaLab/questack-api/pkg/presenter"
	"github.com/HazeyamaLab/questack-api/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type OwnerController interface {
	Create(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type ownerController struct {
	ownerService service.OwnerService
}

func NewOwnerController(s service.OwnerService) OwnerController{
	return &ownerController{s}
}

func (o *ownerController)Create(ctx *gin.Context){
	var ownerForm presenter.OwnerCreateForm
	err := ctx.BindJSON(&ownerForm)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	id, err :=  o.ownerService.Create(ownerForm.Name, ownerForm.Mail, ownerForm.Password)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	token, err := util.GenerateOwnerJwt(id)

	if err != nil{
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "Authorization": token})

}

func (o *ownerController)SignIn(ctx *gin.Context){
	var form presenter.OwnerSignInForm
	err := ctx.BindJSON(&form)
	if err != nil{
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ownerID, err := o.ownerService.SignIn(form.Mail, form.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	token, err := util.GenerateOwnerJwt(ownerID)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"token": token,
	})
}