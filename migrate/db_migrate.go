package main

import (
	"driver_app/initializers"
	"driver_app/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectTODB()
}

func main() {
	db := initializers.DB
	initializers.DB.AutoMigrate(&(models.Role{}), &(models.User{}))
	role := []models.Role{{ID: 1, Name: "driver"}, {ID: 2, Name: "user"}}
	result := db.Create(&role)
	if result.Error != nil {
		log.Fatal("problem ehile runing migration ")
	}
}
