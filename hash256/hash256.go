package hash256

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func (h hasher) Write(s string) (int, error) {
	byts := []byte(s)

	return h.hash.Write(byts)
}

func (h hasher) GetHex() string {
	hashByte := h.hash.Sum([]byte{})
	return hex.EncodeToString(hashByte)
}

func newHasher() hasher {
	return hasher{sha256.New()}
}
