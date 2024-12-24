package decrypting

import (
	"math/big"
)

func decrypt(c, d, n *big.Int) *big.Int {
	var decryptMess big.Int

	decryptMess.Exp(c, d, n)

	return &decryptMess
}
