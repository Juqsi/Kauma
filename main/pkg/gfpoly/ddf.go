package gfpoly

import (
	"math/big"
)

type GfpolyDdf struct {
	F       []string `json:"F"`
	Factors FactorsModelWithDegree
}

func (g *GfpolyDdf) Execute() {
	polyA := NewPolyFromBase64(g.F)
	factors := polyA.Ddf()
	g.Factors = factors.Sort().Base64Degree()

}

// TODO change Factor to FactorModelWithDegree
func (p *Poly) Ddf() Factors {
	X := NewPolyFromBase64([]string{"AAAAAAAAAAAAAAAAAAAAAA==", "gAAAAAAAAAAAAAAAAAAAAA=="})
	exp := big.NewInt(1)

	z := Factors{}
	d := 1
	degree := p.Degree()

	h := new(Poly)
	g := new(Poly)

	for degree >= 2*d {
		exp.Lsh(exp, 128)
		h.PowMod(X, exp, p)
		h.Add(h, X)
		g.Gcd(h, p)

		if !g.IsOne() {
			z = append(z, Factor{*g, d})
			p, _ = p.Div(p, g)
			degree = p.Degree()
		}
		d++
	}

	if !p.IsOne() {
		z = append(z, Factor{*p, p.Degree()})
	} else if len(z) == 0 {
		z = append(z, Factor{*p, 1})
	}

	return z
}
