package gfpoly

type GfpolyGgt struct {
	A []string `json:"A"`
	B []string `json:"B"`
	G []string `json:"G"`
}

func (g *GfpolyGgt) Execute() {
	polyA := NewPolyFromBase64(g.A)
	polyB := NewPolyFromBase64(g.B)
	polyA.Gcd(polyA, polyB)
	g.G = polyA.Base64()
}

func (p *Poly) Gcd(a, b *Poly) *Poly {
	if a.IsZero() {
		p.makeMonic(b)
		return p
	}

	if b.IsZero() {
		p.makeMonic(a)
		return p
	}

	if a.Cmp(b) == -1 {
		return p.Gcd(b, a)
	}
	_, remainder := a.Div(a, b)
	if !remainder.IsZero() {
		return p.Gcd(b, remainder)
	}
	return p.makeMonic(b)
}
