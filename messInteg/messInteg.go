package messinteg

import (
	"crypto/sha256"
	"encoding/hex"
)

func checksumMatches(message string, checksum string) bool {
	var h = sha256.New()
	h.Write([]byte(message))

	var messHash = hex.EncodeToString(h.Sum([]byte{}))

	return messHash == checksum
}
