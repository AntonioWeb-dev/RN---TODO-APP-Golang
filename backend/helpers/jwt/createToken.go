package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id, secret string) (string, error) {
	signSecret := []byte(secret)

	type MyCustomClaims struct {
		UserId string `json:"userId"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signSecret)
}
