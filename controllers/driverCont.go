package controllers

import (
	"driver_app/models"
	"net/http"

	// "gopkg.in/go-playground/validator.v9"
	"github.com/gin-gonic/gin"
)

func (env *Env) CreateDriver(c *gin.Context) {
	//read data from Context
	// log.Fatal("this is connected ////////////////////")
	var data models.SignUp
	err := c.ShouldBindJSON(&data)

	data.RoleID = 1

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"user": data,
			"err":  err.Error(), // cast it to string before showing,
		})
		return

	}

	res, err := env.User.CreateUser(data)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"user": data,
			"err":  err.Error(), // cast it to string before showing,
		})
		return

	}

	c.JSON(http.StatusAccepted, gin.H{
		"code": http.StatusAccepted,
		"user": res,
		"err":  err, // cast it to string before showing,
	})
}
