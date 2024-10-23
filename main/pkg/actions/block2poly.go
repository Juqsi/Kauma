package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

/*
func Block2poly(block string) (ret []uint) {
	a := utils.NewLongFromBase64(block).BigInt()
	return Number2Coefficients(a)
}
*/

type Block2Poly struct {
	Semantic string `json:"semantic"`
	Block    string `json:"block"`
	Result   []uint `json:"coefficients"`
}

func (b *Block2Poly) Execute() {
	b.Result = Number2Coefficients(utils.NewLongFromBase64(b.Block).BigInt())
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
