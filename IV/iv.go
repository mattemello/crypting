package iv

import (
	"math/rand"
)

var count = 0

func generateIV(length int) ([]byte, error) {
	iv := make([]byte, length)
	for i := 0; i < length; i++ {
		count, err := rand.Read([]byte{byte(i)})
		if err != nil {
			return nil, err
		}

		iv[i] = byte(count)
	}
	return iv, nil
}
