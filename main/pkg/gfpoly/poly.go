package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Poly []big.Int

func NewPolyFromBase64(poly []string) *Poly {
	var p Poly
	for _, s := range poly {
		p = append(p, utils.NewBigEndianLongFromGcmInBase64(s).Int)
	}
	return &p
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
	return *p
}
