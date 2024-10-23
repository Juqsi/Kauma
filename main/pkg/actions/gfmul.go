package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Gfmul struct {
	Semantic string `json:"Semantic"`
	Factor1  string `json:"a"`
	Factor2  string `json:"b"`
	Result   string `json:"product"`
}

func (g *Gfmul) Execute() {
	result := GfmulBigInt(utils.NewLongFromBase64(g.Factor1).BigInt(), utils.NewLongFromBase64(g.Factor2).BigInt(), Coeff2Number([]uint{128, 7, 2, 1, 0}))
	g.Result = utils.NewLongFromBigInt(result).GetBase64(16)
}

func GfmulBigInt(factor1, factor2, reduce *big.Int) *big.Int {
	result := big.NewInt(0)
	tmpFactor1 := new(big.Int)
	tmpFactor2 := new(big.Int)

	if factor1.BitLen() < factor2.BitLen() {
		tmpFactor1.Set(factor2)
		tmpFactor2.Set(factor1)
	} else {
		tmpFactor1.Set(factor1)
		tmpFactor2.Set(factor2)
	}

	tmpReduce := new(big.Int).Set(reduce)

	for tmpFactor2.BitLen() > 0 {
		if tmpFactor2.Bit(0) == 1 {
			result.Xor(result, tmpFactor1)
		}

		tmpFactor1.Lsh(tmpFactor1, 1)

		if tmpFactor1.BitLen() >= reduce.BitLen() {
			tmpFactor1.Xor(tmpFactor1, tmpReduce)
		}

		tmpFactor2.Rsh(tmpFactor2, 1)
	}
	return result
}
