package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"math/big"
)

type GfpolySqrt struct {
	Q []string `json:"Q"`
	S []string `json:"S"`
}

func (g *GfpolySqrt) Execute() {
	polyA := NewPolyFromBase64(g.Q)
	polyA.GfSqrt128(polyA)
	g.S = polyA.Base64()
}

func (p *Poly) GfSqrt128(q *Poly) *Poly {
	return p.Sqrt(q, actions.Coeff2Number([]uint{128, 7, 2, 1, 0}))
}

func (p *Poly) Sqrt(q *Poly, m big.Int) *Poly {
	exp := new(big.Int).Lsh(big.NewInt(1), uint(m.BitLen()-2))
	sqrt := make(Poly, (len(*q)+1)/2)
	for i := 0; i < len(*q); i += 2 {
		sqrt[i/2] = actions.Pow(&(*q)[i], exp)
	}
	*p = *sqrt.CutLeadingZeroFaktors()
	return p
}
