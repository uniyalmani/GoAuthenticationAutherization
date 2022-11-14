package routes

import (
	controllers "driver_app/controllers"
	initializers "driver_app/initializers"
	"driver_app/middleware"
	models "driver_app/models"

	"github.com/gin-gonic/gin"
)

func DriverRoute(incomingRoutes *gin.Engine) {
	env := &controllers.Env{
		User: models.UserDbModel{DB: initializers.DB},
	}
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/driver/data", env.GetAllUsers)
}
