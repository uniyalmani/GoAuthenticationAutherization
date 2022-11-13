package main

import (
	"driver_app/controllers"
	"driver_app/initializers"
	"driver_app/models"

	"github.com/gin-gonic/gin"
	// _ "github.com/joho/godotenv/autoload"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()

}

func main() {
	router := gin.Default()

	env := &controllers.Env{
		User: models.UserDbModel{DB: initializers.DB},
	}

	router.POST("/driver/signup", env.CreateDriver)
	router.POST("/user/signup", controllers.CreateUser)

	router.Run()
}
