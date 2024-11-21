package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type GfpolyPowmod struct {
	A []string `json:"A"`
	M []string `json:"M"`
	K int      `json:"K"`
	Z []string `json:"p"`
}

func (g *GfpolyPowmod) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyM := NewPolyFromBase64(g.M)
	polyA.Powmod(*polyA, *polyM, g.K)
	g.Z = polyA.Base64()
}

func (p *Poly) Powmod(a, m Poly, k int) Poly {
	var result Poly
	if k == 0 {
		result = Poly{utils.NewLongFromBigInt(*big.NewInt(1)).Int}
		*p = result
		return result
	}
	result = make(Poly, len(m))
	result[0] = utils.NewLongFromBigInt(*big.NewInt(1)).Int
	for k > 0 {
		if k&1 == 1 {
			result.Mul(result, a)
			_, result = new(Poly).Div(result, m)
		}
		a.Mul(a, a)
		_, a = new(Poly).Div(a, m)
		k >>= 1
	}
	*p = result.Reduce()
	return *p
}
