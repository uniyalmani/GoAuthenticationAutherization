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
