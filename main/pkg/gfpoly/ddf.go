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

	X := NewPolyFromBase64([]string{"AAAAAAAAAAAAAAAAAAAAAA==", "gAAAAAAAAAAAAAAAAAAAAA=="})
	exp := big.NewInt(1)

	z := Factors{}
	d := 1
	fStar := p.DeepCopy()
	degree := fStar.Degree()

	for degree >= 2*d {
		exp.Mul(exp, q)
		h := new(Poly).PowMod(X, exp, fStar)
		h.Add(h, X)
		g := new(Poly).Gcd(h, fStar)

		if !g.IsOne() {
			z = append(z, Factor{*g, d})
			fStar, _ = fStar.Div(fStar, g)
			degree = fStar.Degree()
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
