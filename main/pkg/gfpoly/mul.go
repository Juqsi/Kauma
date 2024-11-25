package gfpoly

import (
	"Abgabe/main/pkg/actions"
)

type GfpolyMul struct {
	A []string `json:"A"`
	B []string `json:"B"`
	P []string `json:"p"`
}

func (g *GfpolyMul) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyB := NewPolyFromBase64(g.B)
	polyA.Mul(*polyA, *polyB)
	g.P = polyA.Base64()
}

func (p *Poly) Mul(a, b Poly) Poly {
	result := make(Poly, len(a)+len(b)-1)
	for indexA, factorA := range a {
		for indexB, factorB := range b {
			c := actions.Gfmul128(factorA, factorB)
			result[indexA+indexB].Xor(&result[indexA+indexB], &c)
		}
	}
	*p = result.CutLeadingZeroFaktors()
	return *p
}
