package blockcipher

import (
	"crypto/aes"
	"crypto/des"
	"errors"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	switch getCipherTypeName(cipherType) {

	case "AES":
		aes, err := aes.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
		return aes.BlockSize(), nil
	case "DES":
		des, err := des.NewCipher(make([]byte, keyLen))
		if err != nil {
			return 0, err
		}
		return des.BlockSize(), nil
	default:
		return 0, errors.New("invalid chiper type")
	}
}
