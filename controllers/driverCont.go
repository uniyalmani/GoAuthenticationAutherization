package controllers

import (
	helpers "driver_app/helpers"
	models "driver_app/models"
	"errors"
	"net/http"

	// "gopkg.in/go-playground/validator.v9"
	"github.com/gin-gonic/gin"
)

func (env *Env) CreateDriver(c *gin.Context) {
	//read data from Context
	// log.Fatal("this is connected ////////////////////")
	var data models.SignUp

	if err := c.ShouldBindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(), // cast it to string before showing,
		})
		return

	}
	data.RoleID = 1
	hash_password, err := helpers.HashPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"err":  err.Error(), // cast it to string before showing,
		})
		return
	}
	data.Password = hash_password
	token, refreshtoken, _ := helpers.GenrateAllToken(data.Email, data.RoleID)
	userId, err := env.User.CreateUser(data)
	data.Token = token
	data.RefreshToken = refreshtoken

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"err":  err.Error(), // cast it to string before showing,
		})
		return

	}

	c.JSON(http.StatusAccepted, gin.H{
		"code":   http.StatusAccepted,
		"userid": userId,
		"user":   data,
		"err":    err, // cast it to string before showing,
	})
}

func (env *Env) LoginDriver(c *gin.Context) {
	var data models.SignIn

	if err := c.ShouldBindJSON(&data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(), // cast it to string before showing,
		})
		return
	}
	data.RoleID = 1
	user := env.User.GetUser(data)

	if !helpers.CheckPasswordHash(data.Password, user.Password) {
		new_err := errors.New("wrong email or password")
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": new_err.Error(), // cast it to string before showing,
		})
		return
	}

	token, refreshtoken, _ := helpers.GenrateAllToken(data.Email, data.RoleID)

	user.Token = token
	user.RefreshToken = refreshtoken

	c.JSON(http.StatusAccepted, gin.H{
		"code": http.StatusAccepted,
		"user": user,
	})
}

func (env *Env) GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"email": c.GetString("email")})

}
