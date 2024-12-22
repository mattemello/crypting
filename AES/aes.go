package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return cipher.Open(nil, nonce, ciphertext, nil)
}
