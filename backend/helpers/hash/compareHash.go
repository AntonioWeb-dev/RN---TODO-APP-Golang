package hash

import "golang.org/x/crypto/bcrypt"

func CompareHash(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
