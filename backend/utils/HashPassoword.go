package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("The error in hashing password from the given password is : ", err)
	}
	return string(hashedPassword)
}

func VerifyHashPassword(password string, hashedPassword string) error {
	compare := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return compare
}
