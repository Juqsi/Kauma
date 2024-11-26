package gfpoly

type GfpolyAdd struct {
	A []string `json:"A"`
	B []string `json:"B"`
	Z []string `json:"Z"`
}

func (g *GfpolyAdd) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyB := NewPolyFromBase64(g.B)
	polyA.Add(polyA, polyB)
	g.Z = polyA.Base64()
}

func (p *Poly) Add(a, b *Poly) (result *Poly) {
	if len(*a) > len(*b) {
		a, b = b, a
	}
	for i := 0; i < len(*a); i++ {
		(*b)[i].Xor(&(*a)[i], &(*b)[i])
	}
	*p = *b.CutLeadingZeroFaktors()
	return p
}
