package server

import (
	"github.com/HazeyamaLab/questack-api/conf"
	"github.com/HazeyamaLab/questack-api/pkg/controller"
	"github.com/HazeyamaLab/questack-api/pkg/domain/service"
	"github.com/HazeyamaLab/questack-api/pkg/middleware"
	"github.com/HazeyamaLab/questack-api/pkg/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(){
	r := router()
	r.Run(":8888")
}

func router() *gin.Engine{
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AddAllowHeaders("Authorization")

	r.Use(cors.New(config))
	r.Use(gin.Logger())

	ownerCtr := ownerInjector()
	stackCtr := stackInjector()
	questionCtr := questionInjector()

	api := r.Group("/questack-api")
	api.GET("/stacks/:token", stackCtr.FindOneByURL)
	api.POST("/sign-in", ownerCtr.SignIn)
	api.POST("/user", ownerCtr.Create)
	api.POST("/question/create",questionCtr.Create)
	auth := api.Use(middleware.Auth())
	auth.GET("/stacks", stackCtr.FindMine)
	auth.POST("/stack", stackCtr.Create)
	auth.POST("/stack/update", stackCtr.Update)
	return r
}

func ownerInjector() controller.OwnerController{
	db := conf.GetDB()
	ownerRepository := repository.NewOwnerRepository(db)
	ownerService := service.NewOwnerService(ownerRepository)
	return controller.NewOwnerController(ownerService)
}

func stackInjector() controller.StackController{
	db := conf.GetDB()
	stackRepository := repository.NewStackRepository(db)
	stackService := service.NewStackService(stackRepository)
	return controller.NewStackController(stackService)
}

func questionInjector() controller.QuestionController{
	db := conf.GetDB()
	questionRepository := repository.NewQuestionRepository(db)
	stackRepository := repository.NewStackRepository(db)
	questionService := service.NewQuestionService(questionRepository, stackRepository)
	return controller.NewQuestionController(questionService)
}
