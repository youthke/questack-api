package server

import (
	"github.com/HazeyamaLab/questack-api/conf"
	"github.com/HazeyamaLab/questack-api/pkg/controller"
	"github.com/HazeyamaLab/questack-api/pkg/domain/service"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
	"github.com/gin-gonic/gin"
)

func Init(){
	r := router()
	r.Run(":8888")
}

func router() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())

	ctr := ownerInjector()

	api := r.Group("/questack-api")
	api.POST("/user", ctr.Create)

	return r
}

func ownerInjector() controller.OwnerController{
	db := conf.GetDB()
	ownerRepository := repository.NewOwnerRepository(db)
	ownerService := service.NewOwnerService(ownerRepository)
	ownerController := controller.NewOwnerController(ownerService)
	return ownerController
}
