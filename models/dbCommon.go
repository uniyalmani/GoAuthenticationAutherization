package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   int
	Name string
}

type User struct {
	gorm.Model
	ID        int
	FirstName string
	LastName  string
	Password  string
	Email     string
	RoleID    int
	Role      Role
}

// attache all methods related to Users table
type UserDbModel struct {
	DB *gorm.DB
}

func (u UserDbModel) CreateUser(data SignUp) (int, error) {
	user := User{FirstName: data.FirstName, LastName: data.LastName, Password: data.Password,
		RoleID: data.RoleID, Email: data.Email}
	res := u.DB.Create(&user)
	return user.ID, res.Error
}

func (u UserDbModel) GetUser(data SignIn) SignUp {
	var result SignUp
	u.DB.Table("users").Select("*").Where("email = ?", data.Email).Scan(&result)

	return result
}
