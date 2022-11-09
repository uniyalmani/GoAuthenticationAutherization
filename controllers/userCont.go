package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Print("hello inside user")
}
