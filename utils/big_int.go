package utils

import "math/big"

func MinBigInt(l, r *big.Int) *big.Int {
	if l.Cmp(r) <= 0 {
		return l
	}
	return r
}
