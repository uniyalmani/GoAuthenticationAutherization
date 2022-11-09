package main

import (
	"github.com/gin-gonic/gin"
	"driver_app/initializers"
	"driver_app/controllers"

	// _ "github.com/joho/godotenv/autoload"
)




func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
}

func main() {
	router := gin.Default()
	
	router.POST("/driver/signup", controllers.CreateDriver)
	router.POST("/user/signup", controllers.CreateUser)

	router.Run()
}
