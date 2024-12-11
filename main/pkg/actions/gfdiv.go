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
	return GfdivBigInt(a, b, Reduce128)
}

func GfdivBigInt(a, b, reduce big.Int) big.Int {
	bInverse := GfInverseBigInt(b, reduce)
	return GfmulBigInt(a, bInverse, reduce)
}

func GfInverseBigInt(b, reduce big.Int) big.Int {
	var u, v, g1, g2 big.Int

	u.Set(&b)
	v.Set(&reduce)
	g1.SetInt64(1)
	g2.SetInt64(0)

	for v.Sign() != 0 {
		degU := u.BitLen() - 1
		degV := v.BitLen() - 1
		shift := degU - degV

		if shift >= 0 {
			// Reduktion von u
			tempV := new(big.Int).Lsh(&v, uint(shift)) // v * x^shift
			u.Xor(&u, tempV)                           // u = u - v * x^shift (XOR)

			// Reduktion von g1
			tempG2 := new(big.Int).Lsh(&g2, uint(shift))
			g1.Xor(&g1, tempG2) // g1 = g1 - g2 * x^shift (XOR)
		}

		u, v = v, u
		g1, g2 = g2, g1
	}

	inverseLen := g1.BitLen()
	reverseLen := reduce.BitLen()
	if inverseLen >= reverseLen {
		g1.Xor(&g1, reduce.Lsh(&reduce, uint(inverseLen-reverseLen)))
	}
	return g1
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
