package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yermanovberik/golang-final-project/internal/controller"
	"github.com/yermanovberik/golang-final-project/pkg"
)

func main() {
	fmt.Println(os.Getenv("PORT"))
	app, err := pkg.App()

	if err != nil {
		log.Fatal(err)
	}
	ginRouter := gin.Default()
	controller.Setup(app, ginRouter)
	
	ginRouter.Run(fmt.Sprintf(":%s", app.Env.PORT))
}
