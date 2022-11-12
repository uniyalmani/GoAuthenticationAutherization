package main

import (
	"driver_app/controllers"
	"driver_app/initializers"

	"github.com/gin-gonic/gin"
	// _ "github.com/joho/godotenv/autoload"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
}

func main() {
	router := gin.Default()

	router.GET("/driver/signup", controllers.CreateDriver)
	router.POST("/user/signup", controllers.CreateUser)

	router.Run()
}
