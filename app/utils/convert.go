package utils

import (
	"math/big"
)

func ConvertWeiToEth(v *big.Int) *big.Rat {
	if v == nil {
		return big.NewRat(0, 1)
	}

	return big.NewRat(v.Int64(), 1000000000000000000)
}
