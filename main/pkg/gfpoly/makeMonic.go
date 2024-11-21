package gfpoly

import "Abgabe/main/pkg/actions"

type GfpolyMakeMonic struct {
	A          []string `json:"A"`
	ASternchen []string `json:"A*"`
}

func (g *GfpolyMakeMonic) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyA.makeMonic()
	g.ASternchen = polyA.Base64()
}

func (poly *Poly) makeMonic() Poly {
	l := len(*poly)
	for i := 0; i < l; i++ {
		(*poly)[i] = actions.Gfdiv128((*poly)[i], (*poly)[l-1])
	}
	return *poly
}
