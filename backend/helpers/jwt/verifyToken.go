package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// VerifyToken - make the token's validation
func VerifyToken(token, secret string) (string, error) {
	signSecret := []byte(secret)
	t, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Sign method error")
			}
			return signSecret, nil
		},
	)
	if err != nil {
		return "", err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		userID := fmt.Sprintf("%v", claims["userId"])
		return userID, nil
	}
	return "", errors.New("jwt invalid")
}
