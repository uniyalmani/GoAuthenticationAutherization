package main

import (
	initializers "driver_app/initializers"
	routes "driver_app/routes"

	"github.com/gin-gonic/gin"
	// _ "github.com/joho/godotenv/autoload"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()

}

func main() {
	router := gin.Default()

	routes.AuthRoutes(router)
	routes.DriverRoute(router)
	router.Run()
}
