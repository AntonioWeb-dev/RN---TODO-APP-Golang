package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// VerifyToken - make the token's validation
func VerifyToken(token, secret string) error {
	t, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Sign method error")
			}
			return []byte(secret), nil
		},
	)
	if err != nil {
		return err
	}
	if _, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return nil
	}
	return errors.New("jwt invalid")
}
