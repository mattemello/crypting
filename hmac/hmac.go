package hmac

import (
	"crypto/sha256"
	"fmt"
)

func hmac(message, key string) string {
	firstHalf := ""
	secondHalf := ""

	firstHalf = key[len(key)/2:]
	secondHalf = key[:len(key)/2]

	h := sha256.New()
	h2 := sha256.New()

	h2.Write(append([]byte(firstHalf), []byte(message)...))

	h.Write(append([]byte(secondHalf), h2.Sum(nil)...))

	return fmt.Sprintf("%x", h.Sum(nil))

}
