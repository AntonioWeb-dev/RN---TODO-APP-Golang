package hash

import "golang.org/x/crypto/bcrypt"

func GenerateHash(data string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}
