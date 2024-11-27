package gfpoly

type GfpolyDiff struct {
	F       []string `json:"F"`
	FStrich []string `json:"F'"`
}

func (g *GfpolyDiff) Execute() {
	polyA := NewPolyFromBase64(g.F)
	polyFStrich := new(Poly)
	polyFStrich.Diff(polyA)
	g.FStrich = polyFStrich.Base64()
}

func (p *Poly) Diff(a *Poly) *Poly {
	result := make(Poly, len(*a))
	for i := 1; i < len(*a); i += 2 {
		result[i-1] = (*a)[i]
	}
	result.CutLeadingZeroFaktors()
	*p = result
	return &result
}
