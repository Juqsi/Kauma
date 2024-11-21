package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Poly []big.Int
type Polys []Poly

func NewPolyFromBase64(poly []string) *Poly {
	p := make(Poly, len(poly))
	for i, s := range poly {
		p[i] = utils.NewBigEndianLongFromGcmInBase64(s).Int
	}
	return &p
}

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

func (p *Poly) Base64() []string {
	s := make([]string, len(*p))
	for i, num := range *p {
		s[i] = utils.NewLongFromBigInt(num).GcmToggle().GetBase64(16)
	}
	return s
}

func (p *Poly) Reduce() Poly {
	lenP := len(*p)
	for i := lenP; i > 0; i-- {
		if (*p)[i-1].Sign() != 0 {
			*p = (*p)[:i]
			return *p
		}
	}
	*p = Poly{utils.NewLongFromBigInt(*big.NewInt(0)).Int}
	return *p
}

func (p *Poly) Degree() int {
	return len(*p)
}
