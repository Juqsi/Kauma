package gfpoly

import (
	"math/big"
)

type GfpolyDiff struct {
	F       []string `json:"Q"`
	FStrich []string `json:"S"`
}

func (g *GfpolyDiff) Execute() {
	polyA := NewPolyFromBase64(g.F)
	polyFStrich := new(Poly)
	polyFStrich.Diff(polyA)
	g.FStrich = polyFStrich.Base64()
}

func (p *Poly) Diff(a *Poly) *Poly {
	zero := big.NewInt(0)
	*p = *a.DeepCopy()
	for i := range *p {
		if i&1 == 0 {
			(*p)[i] = *zero
		}
	}
	*p = append(*p, *zero)
	*p = (*p)[1:]
	p.CutLeadingZeroFaktors()
	return p
}
