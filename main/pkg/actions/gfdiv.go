package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Gfdiv struct {
	Factor1 string `json:"a"`
	Factor2 string `json:"b"`
	Result  string `json:"q"`
}

func (g *Gfdiv) Execute() {
	factor1 := utils.NewBigEndianLongFromGcmInBase64(g.Factor1).Int
	factor2 := utils.NewBigEndianLongFromGcmInBase64(g.Factor2).Int

	result := Gfdiv128(factor1, factor2)
	g.Result = utils.NewLongFromBigInt(result).GcmToggle().GetBase64(16)
}

func Gfdiv128(a, b big.Int) big.Int {
	return GfdivBigInt(a, b, Coeff2Number([]uint{128, 7, 2, 1, 0}))
}

func GfdivBigInt(a, b, reduce big.Int) big.Int {
	// Inverse of b -> 2^128-2
	exp := *new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 128), big.NewInt(2))
	return GfmulBigInt(a, Pow(&b, &exp), reduce)
}

func Pow(a, exponent *big.Int) big.Int {
	result := *big.NewInt(1)
	base := *new(big.Int).Set(a)
	exp := *new(big.Int).Set(exponent)

	// Square-and-Multiply-Algorithmus
	for exp.Sign() > 0 {
		if exp.Bit(0) == 1 {
			result = Gfmul128(result, base)
		}
		base = Gfmul128(base, base)
		exp.Rsh(&exp, 1)
	}

	return result
}
