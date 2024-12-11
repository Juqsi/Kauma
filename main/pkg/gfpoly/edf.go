package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"math/big"
)

type GfpolyEdf struct {
	F       []string `json:"F"`
	D       int      `json:"d"`
	Factors [][]string
}

func (g *GfpolyEdf) Execute() {
	polyA := NewPolyFromBase64(g.F)
	factors := new(Polys).Edf(polyA, g.D)
	g.Factors = factors.Sort().Base64()

}

func (p *Polys) Edf(f *Poly, d int) Polys {
	q := big.NewInt(1)
	q.Lsh(q, 128)
	n := f.Degree() / d

	z := Polys{*f}

	exp := new(big.Int).Exp(q, big.NewInt(int64(d)), nil)
	exp.Sub(exp, big.NewInt(1))
	exp.Div(exp, big.NewInt(3))

	for len(z) < n {
		h := RandomPolynomial(f.Degree() - 1)

		g := new(Poly).PowMod(h, exp, f)
		g.Add(g, &Poly{actions.OneBlock})

		for i := len(z) - 1; i >= 0; i-- {
			u := z[i]
			if u.Degree() > d {
				j := new(Poly).Gcd(&u, g)
				if !j.IsOne() && j.Cmp(&u) != 0 {
					tmp, _ := new(Poly).Div(&u, j)
					z = append(z[:i], append(z[i+1:], []Poly{*j.makeMonic(j), *tmp.makeMonic(tmp)}...)...)
				}
			}
		}

	}
	*p = z
	return z
}
