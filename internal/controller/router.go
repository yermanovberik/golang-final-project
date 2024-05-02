package controller

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	//_ "github.com/yermanovberik/golang-final-project/docs"
	"github.com/yermanovberik/golang-final-project/middleware"
	
	"github.com/yermanovberik/golang-final-project/internal/repository"
	"github.com/yermanovberik/golang-final-project/pkg"
)

func Setup(app pkg.Application, router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	signupController := &SignupController{
		UserRepository: repository.NewUserRepository(app.Pql),
		Env : app.Env,
	}

	router.POST("/signup", signupController.Signup)
	router.POST("/signin" , signupController.Login)
	router.Use(middleware.JWTAuth(app.Env.AccessTokenSecret))
	router.POST("/logout" , signupController.Logout)
}
