package helpers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type TokenDetails struct {
	Email  string
	RoleID int
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenrateAllToken(email string, role_id int) (string, string, error) {
	claims := TokenDetails{
		Email:  email,
		RoleID: role_id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := TokenDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(48)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	RefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return "", "", err
	}
	return token, RefreshToken, nil

}

func ValidateToken(signedToken string) (claims *TokenDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*TokenDetails)

	if !ok {
		msg = fmt.Sprintf("the token is  invalid")
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is exppired")
		msg = err.Error()
		return
	}
	return claims, msg
}

// func UpdateAllToken(signedToken string, signedRefreshToken string, userEmail string, userRole int){

// }
