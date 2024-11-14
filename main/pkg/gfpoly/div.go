package gfpoly

type GfpolyDiv struct {
	A []string `json:"A"`
	K int      `json:"K"`
	Z []string `json:"p"`
}

func (g *GfpolyDiv) Execute() {
	polyA := NewPolyFromBase64(g.A)
	//polyA.Div(*polyA, g.K)
	g.Z = polyA.Base64()
}

func (p *Poly) Div(a, b Poly) Poly {
	return nil

}
