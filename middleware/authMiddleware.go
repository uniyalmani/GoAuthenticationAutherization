package middleware

import (
	"driver_app/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	clientToken := c.Request.Header.Get("token")
	if clientToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not authenticate user"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)

	if err != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	c.Set("email", claims.Email)
	c.Set("role_id", claims.RoleID)
	c.Next()
}
