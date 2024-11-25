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
