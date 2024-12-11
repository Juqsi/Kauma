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

func (p *Poly) Ddf() Factors {
	q := big.NewInt(1)
	q.Lsh(q, 128)
	z := Factors{}
	d := 1
	fStar := p.DeepCopy()

	for fStar.Degree() >= 2*d {
		X := NewPolyFromBase64([]string{"AAAAAAAAAAAAAAAAAAAAAA==", "gAAAAAAAAAAAAAAAAAAAAA=="})
		exp := new(big.Int).Exp(q, big.NewInt(int64(d)), nil)
		h := new(Poly).PowMod(X, *exp, fStar)
		h.Add(h, X)
		g := new(Poly).Gcd(h, fStar)
		if !g.IsOne() {
			z = append(z, Factor{*g, d})
			fStar, _ = fStar.Div(fStar, g)
		}
		d++
	}

	if !fStar.IsOne() {
		z = append(z, Factor{*fStar, fStar.Degree()})
	} else if len(z) == 0 {
		z = append(z, Factor{*p, 1})
	}

	return z
}
