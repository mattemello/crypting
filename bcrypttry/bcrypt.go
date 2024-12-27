package bcrypttry

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bcry, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", nil
	}

	return string(bcry), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
