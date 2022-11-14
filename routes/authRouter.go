package routes

import (
	controllers "driver_app/controllers"
	initializers "driver_app/initializers"
	models "driver_app/models"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	env := &controllers.Env{
		User: models.UserDbModel{DB: initializers.DB},
	}

	incomingRoutes.POST("/driver/signup", env.CreateDriver)
	incomingRoutes.POST("/driver/login", env.LoginDriver)
}
