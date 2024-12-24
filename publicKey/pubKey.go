package publickey

import (
	"math/big"
)

func encrypt(m, e, n *big.Int) *big.Int {
	var cryptMess big.Int

	cryptMess.Exp(m, e, n)

	return &cryptMess
}
