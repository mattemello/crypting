package des

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(plaintext)%block.BlockSize() != 0 {
		plaintext = padMsg(plaintext, block.BlockSize())
	}

	ciphertext := make([]byte, block.BlockSize()+len(plaintext))
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	desBlock := cipher.NewCBCEncrypter(block, iv)
	desBlock.CryptBlocks(ciphertext[block.BlockSize():], plaintext)

	return ciphertext, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	for i := 0; i < len(plaintext); i += blockSize {
		if i+blockSize > len(plaintext) {
			newBlock := padWithZeros(plaintext[i:], blockSize)
			plaintext = append(plaintext[:i], newBlock...)
			break
		}
	}

	return plaintext
}
