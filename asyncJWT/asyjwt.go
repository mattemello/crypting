package asyncjwt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	h := sha256.New()

	h.Write([]byte(message))

	signature, err := ecdsa.SignASN1(rand.Reader, privateKey, h.Sum(nil))
	if err != nil {
		return "", err
	}

	signHex := fmt.Sprintf("%x", signature)

	return message + "." + signHex, nil

}
