package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type GfpolyPowmod struct {
	A []string `json:"A"`
	M []string `json:"M"`
	K big.Int  `json:"K"`
	Z []string `json:"p"`
}

func (g *GfpolyPowmod) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyM := NewPolyFromBase64(g.M)
	polyA.Powmod(polyA, polyM, g.K)
	g.Z = polyA.Base64()
}

func (p *Poly) Powmod(base *Poly, m *Poly, k big.Int) *Poly {
	var result = new(Poly)
	base = base.DeepCopy()

	if k.Sign() == 0 {
		result = &Poly{utils.NewLongFromBigInt(*big.NewInt(1)).Int}
		*p = *result
		return result
	}
	*result = make(Poly, len(*m))
	(*result)[0] = utils.NewLongFromBigInt(*big.NewInt(1)).Int
	for k.Sign() != 0 {
		if k.Bit(0)&1 == 1 {
			result.Mul(result, base)
			_, result = new(Poly).Div(result, m)
		}
		base.Mul(base, base)
		_, base = new(Poly).Div(base, m)
		k.Rsh(&k, 1)
	}
	*p = *result.CutLeadingZeroFaktors()
	return p
}
