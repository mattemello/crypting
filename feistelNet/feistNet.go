package feistelnet

import "fmt"

func feistel(msg []byte, roundeKey [][]byte) []byte {
	halfMsgDim := len(msg) / 2

	rhs := msg[:halfMsgDim]
	lhs := msg[halfMsgDim:]

	fmt.Println(rhs, lhs, msg)

	for _, rounKey := range roundeKey {
		nextRhs := xor(rhs, hash(rhs, rounKey, len(rounKey)))
		lhs = nextRhs
		rhs = lhs
	}

	return append(rhs, lhs...)
}
