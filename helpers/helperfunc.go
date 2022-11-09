package helpers

import (
	"golang.org/x/crypto/bcrypt"
)


type Password struct{
}

func (p Password) HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(hash), err
}


func (p Password) CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

