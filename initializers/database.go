package initializers

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDBUrl() string {
	var err error
	envVariables, err := godotenv.Read()

	if err != nil {
		log.Fatal("not found env variable", err)
	}
	DB_USERNAME := envVariables["MYSQL_USER"]
	DB_PASSWORD := envVariables["MYSQL_ROOT_PASSWORD"]
	DB_NAME := envVariables["MYSQL_DATABASE"]
	DB_HOST := envVariables["MYSQL_ROOT_HOST"]
	DB_PORT := envVariables["DB_PORT"]

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	return dsn
}

var DB *gorm.DB

func ConnectTODB() {
	var err error
	dsn := getDBUrl()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error in connecting with db", err, "|||", dsn)
	}

}
