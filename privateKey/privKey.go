package privatekey

import (
	"math/big"
)

// Get the private exponent
func getD(e, tot *big.Int) *big.Int {
	var invMod big.Int

	invMod.ModInverse(e, tot)

	return &invMod
}
