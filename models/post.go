package models

// import (
// 	//"gopkg.in/validator.v2"
// 	"github.com/gin-gonic/gin"
// )

type SignUp struct {
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	RoleID       int    `json:"roleid"`
	Token        string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
}

type SignIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	RoleID   int
}
