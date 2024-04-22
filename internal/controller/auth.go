package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yermanovberik/golang-final-project/internal/models"
)

type SignupController struct {
	UserRepository models.UserRepository
}


func (sc *SignupController) Login(c *gin.Context){
	fmt.Println("qweqwe")
}