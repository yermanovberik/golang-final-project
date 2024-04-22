package controller

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/yermanovberik/golang-final-project/docs"

	"github.com/yermanovberik/golang-final-project/pkg"
)

func Setup(app pkg.Application, router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Static("/images", "./images")

}
