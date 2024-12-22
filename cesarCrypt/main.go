package main

import (
	"strings"
)

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, (-1 * key))
}

func crypt(text string, key int) string {
	crypted := ""

	for _, char := range text {
		crypted += getOffsetChar(char, key)
	}

	return crypted
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	positionC := strings.Index(alphabet, string(c))

	positionToReturn := positionC + offset

	if positionToReturn < 0 {
		positionToReturn = len(alphabet) + positionToReturn
	}

	if positionToReturn > len(alphabet) {
		positionToReturn -= len(alphabet)
	}

	return string(alphabet[positionToReturn])
}
