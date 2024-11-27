package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
	"sort"
)

type GfpolySff struct {
	F       []string `json:"F"`
	Factors FactorsModel
}
type Factor struct {
	Factor   Poly
	Exponent int
}

type Factors []Factor

type FactorModel struct {
	Factor   []string `json:"factor"`
	Exponent int      `json:"exponent"`
}
type FactorsModel []FactorModel

func (list Factors) Base64() FactorsModel {
	factorsModel := make(FactorsModel, len(list))
	for i, f := range list {
		factorsModel[i].Factor = f.Factor.Base64()
		factorsModel[i].Exponent = f.Exponent
	}
	return factorsModel
}

func (g *GfpolySff) Execute() {
	polyA := NewPolyFromBase64(g.F)
	factors := polyA.sff()
	g.Factors = factors.Sort().Base64()
}

func (f *Poly) sff() Factors {
	one := Poly{utils.NewLongFromBigInt(*big.NewInt(1)).Int}
	c := new(Poly).Gcd(f, new(Poly).Diff(f))
	f, _ = f.Div(f, c)
	z := new(Factors)
	e := 1

	for f.Cmp(f, &one) != 0 {
		y := new(Poly).Gcd(f, c)
		if f.Cmp(f, y) != 0 {
			t, _ := f.Div(f, y)
			*z = append(*z, Factor{*t, e})
		}
		*f = *y
		c, _ = c.Div(c, y)
		e++
	}

	if c.Cmp(c, &one) != 0 {
		r := new(Poly).GfSqrt128(c).sff()
		for _, factor := range r {
			*z = append(*z, Factor{factor.Factor, 2 * factor.Exponent})
		}
	}

	return *z
}

func (list *Factors) Sort() *Factors {
	sort.SliceStable(*list, func(i, j int) bool {
		return (*list)[i].Factor.Cmp(&(*list)[i].Factor, &(*list)[j].Factor) < 0
	})
	return list
}
