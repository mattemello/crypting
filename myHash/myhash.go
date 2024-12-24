package myhash

import "math/bits"

func hash(input []byte) [4]byte {
	var hashTable [4]byte

	for i, bite := range input {
		input[i] = bits.RotateLeft8(bite, 3)
		input[i] <<= 2
		hashTable[i%4] ^= input[i]
	}

	return hashTable
}
