package mac

import (
	"crypto/sha256"
	"fmt"
)

func macMatches(message, key, checksum string) bool {
	h := sha256.New()
	h.Write(append([]byte(message), []byte(key)...))
	return checksum == fmt.Sprintf("%x", h.Sum(nil))
}
