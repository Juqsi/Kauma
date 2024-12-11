package gfpoly

import (
	"Abgabe/main/pkg/actions"
)

type GfpolyPow struct {
	A []string `json:"A"`
	K int      `json:"K"`
	Z []string `json:"p"`
}

func (g *GfpolyPow) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyA.Pow(polyA, g.K)
	g.Z = polyA.Base64()
}

func (p *Poly) Pow(a *Poly, n int) *Poly {
	result := &Poly{actions.OneBlock}
	for n > 0 {

		if n&1 == 1 {
			result.Mul(result, a)
		}
		a.Mul(a, a)
		n = n >> 1
	}
	*p = *result.CutLeadingZeroFaktors()
	return p
}
