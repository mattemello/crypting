package salts

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt(length int) ([]byte, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)

	return salt, err
}

func hashPassword(password, salt []byte) []byte {
	password = append(password, salt...)

	h := sha256.New()

	h.Write(password)

	return h.Sum(nil)
}
