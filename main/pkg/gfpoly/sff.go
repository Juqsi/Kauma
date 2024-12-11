package gfpoly

type GfpolySff struct {
	F       []string `json:"F"`
	Factors FactorsModel
}

func (g *GfpolySff) Execute() {
	polyA := NewPolyFromBase64(g.F)
	factors := polyA.sff()
	g.Factors = factors.Sort().Base64()
}

func (f *Poly) sff() Factors {
	diff := new(Poly).Diff(f)
	c := diff.Gcd(f, diff)
	f, _ = f.Div(f, c)
	z := new(Factors)
	e := 1

	for !f.IsOne() {
		y := new(Poly).Gcd(f, c)
		if f.Cmp(y) != 0 {
			t, _ := f.Div(f, y)
			*z = append(*z, Factor{*t, e})
		}
		*f = *y
		c, _ = c.Div(c, y)
		e++
	}

	if !c.IsOne() {
		sqrtFactors := new(Poly).GfSqrt128(c).sff()
		for _, factor := range sqrtFactors {
			*z = append(*z, Factor{factor.Factor, 2 * factor.Exponent})
		}
	}

	return *z
}
