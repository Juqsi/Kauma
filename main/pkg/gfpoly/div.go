package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/utils"
	"math/big"
)

type GfpolyDiv struct {
	A []string `json:"A"`
	B []string `json:"B"`
	Q []string `json:"Q"`
	R []string `json:"R"`
}

func (g *GfpolyDiv) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyB := NewPolyFromBase64(g.B)
	polyQ, polyR := new(Poly).Div(polyA, polyB)
	g.Q = polyQ.Base64()
	g.R = polyR.Base64()
}

func (p *Poly) Div(dividend, divisor *Poly) (quotient, remainder *Poly) {
	divisor.CutLeadingZeroFaktors()
	dividend.CutLeadingZeroFaktors()
	lenDivisor := len(*divisor)
	if len(*dividend) < len(*divisor) {
		return &Poly{utils.NewLongFromBigInt(*big.NewInt(0)).Int}, dividend
	}
	quotient = new(Poly)
	*quotient = make(Poly, len(*dividend)-lenDivisor+1)
	for len(*dividend) >= lenDivisor {
		lenDividen := len(*dividend)
		tmp := make(Poly, lenDividen-lenDivisor+1)
		a := (*dividend)[lenDividen-1]
		b := (*divisor)[lenDivisor-1]
		tmp[len(tmp)-1] = actions.Gfdiv128(a, b)
		(*quotient)[lenDividen-lenDivisor] = tmp[len(tmp)-1]
		tmp.Mul(&tmp, divisor)
		dividend.Add(dividend, &tmp)
		if len(*dividend) == lenDividen {
			break
		}
	}
	return quotient.CutLeadingZeroFaktors(), dividend.CutLeadingZeroFaktors()
}
