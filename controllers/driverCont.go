package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateDriver(c *gin.Context) {
	//read data from Context
	// log.Fatal("this is connected ////////////////////")
	c.JSON(200, gin.H{
		"code":    200,
		"message": "created", // cast it to string before showing
	})
}
