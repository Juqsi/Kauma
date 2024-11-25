package gfpoly

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Poly []big.Int

func NewPolyFromBase64(poly []string) *Poly {
	p := make(Poly, len(poly))
	for i, s := range poly {
		p[i] = utils.NewBigEndianLongFromGcmInBase64(s).Int
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

// TODO changr to Square multiply
func (p *Poly) CutLeadingZeroFaktors() Poly {
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
	p.CutLeadingZeroFaktors()
	l := len(*p)
	if l == 1 && (*p)[0].Sign() == 0 {
		return -1
	}
	return l - 1
}

func (p *Poly) IsZero() bool {
	return p.Degree() == -1
}

// Cmp compares x and y and returns:
//   - -1 if x < y;
//   - 0 if x == y;
//   - +1 if x > y.
func (p *Poly) Cmp(x, y Poly) int {
	xDegree := x.Degree()
	yDegree := y.Degree()
	if xDegree != yDegree {
		if xDegree > yDegree {
			return 1
		} else {
			return -1
		}
	} else if xDegree == yDegree {
		index := xDegree
		for index >= 0 {
			a := x[index].CmpAbs(&y[index])
			if a != 0 {
				return a
			}
			index--
		}
		return 0
	}
	panic("should not happen Cmp Factor")
}
