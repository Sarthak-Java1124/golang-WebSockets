package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	Email string `json:"email"`
	*jwt.RegisteredClaims
}

var JWT_SECRET = []byte("mynameissarthakharsh")

func GenerateJWT(email string) (string, error) {
	claims := JWTClaims{
		Email: email,
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 10)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(JWT_SECRET)
	if err != nil {
		log.Println("The error in signing the token is :", err)
	}

	return signed, err

}
