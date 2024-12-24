package tote

import (
	"crypto/rand"
	"math/big"
)

func getTot(p, q *big.Int) *big.Int {
	var tot big.Int

	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))
	tot.Mul(p, q)

	return &tot
}

func getE(tot *big.Int) *big.Int {
	var forMod big.Int
	for {
		e, err := rand.Int(randReader, tot)
		if err != nil {
			return nil
		}

		if e.Cmp(big.NewInt(1)) == 1 || forMod.Mod(e, tot).Cmp(big.NewInt(1)) == 0 {
			return e
		}
	}
}
