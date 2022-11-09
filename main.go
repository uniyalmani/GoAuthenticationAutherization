package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	// _ "github.com/joho/godotenv/autoload"
)





func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	
	router.Run()
}
