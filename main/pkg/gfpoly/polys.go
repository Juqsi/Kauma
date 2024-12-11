package gfpoly

type Polys []Poly

func NewPolyListFromBase64(polys [][]string) *Polys {
	p := make(Polys, len(polys))
	for i, poly := range polys {
		p[i] = *NewPolyFromBase64(poly)
	}
	return &p
}

func (polys *Polys) Base64() [][]string {
	result := make([][]string, len(*polys))
	for i, poly := range *polys {
		result[i] = poly.Base64()
	}
	return result
}

func (p *Polys) removePoly(polys *Polys, poly *Poly) *Polys {
	*p = make(Polys, 0, len(*polys))
	for _, pol := range *polys {
		if pol.Cmp(poly) != 0 {
			*p = append(*p, pol)
		}
	}
	return p
}
