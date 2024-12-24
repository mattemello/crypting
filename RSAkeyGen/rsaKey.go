package rsakeygen

import (
	"math/big"
)

func generatePrivateNums(keysize int) (*big.Int, *big.Int) {
	p, err := getBigPrime(keysize)
	if err != nil {
		panic(err)
	}

	q, err := getBigPrime(keysize)
	if err != nil {
		panic(err)
	}

	return p, q
}

// Calculate n = p * q
func getN(p, q *big.Int) *big.Int {
	var n big.Int

	n.Mul(p, q)

	return &n
}
