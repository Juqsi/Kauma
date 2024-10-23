package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

func Block2poly(block string) (ret []uint) {
	a := utils.NewLongFromBase64(block).BigInt()
	return Number2Coefficients(a)
}

func Number2Coefficients(number *big.Int) (ret []uint) {
	bitLen := number.BitLen()
	for i := 0; i < bitLen; i++ {
		if number.Bit(i) == 1 {
			ret = append(ret, uint(i))
		}
	}
	return
}
