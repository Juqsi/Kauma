package gfpoly

import (
	"Abgabe/main/pkg/actions"
)

type GfpolyMakeMonic struct {
	A          []string `json:"A"`
	ASternchen []string `json:"A*"`
}

func (g *GfpolyMakeMonic) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyA.makeMonic(polyA)
	g.ASternchen = polyA.Base64()
}

func (p *Poly) makeMonic(poly *Poly) *Poly {
	poly.CutLeadingZeroFaktors()
	l := len(*poly)
	tmp := make(Poly, l)
	for i := 0; i < l; i++ {
		tmp[i] = actions.Gfdiv128((*poly)[i], (*poly)[l-1])
	}
	*p = tmp
	return p
}
