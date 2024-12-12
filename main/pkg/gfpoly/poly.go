package gfpoly

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/utils"
	"math/big"
	"math/rand"
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

func (p *Poly) CutLeadingZeroFaktors() *Poly {
	lenP := len(*p)
	for i := lenP; i > 0; i-- {
		if (*p)[i-1].Sign() != 0 {
			*p = (*p)[:i]
			return p
		}
	}
	*p = Poly{utils.NewLongFromBigInt(*big.NewInt(0)).Int}
	return p
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

func (p *Poly) IsOne() bool {
	return p.Degree() == 0 && (*p)[0].Cmp(big.NewInt(1)) == 0
}

// Cmp compares x and y and returns:
//   - -1 if x < y;
//   - 0 if x == y;
//   - +1 if x > y.
func (x *Poly) Cmp(y *Poly) int {
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
			a := (*x)[index].CmpAbs(&(*y)[index])
			if a != 0 {
				return a
			}
			index--
		}
		return 0
	}
	panic("should not happen Cmp Factor")
}

func (p *Poly) DeepCopy() *Poly {
	if p == nil {
		return nil
	}
	copy := make(Poly, len(*p))
	for i, v := range *p {
		copy[i] = *big.NewInt(0).Set(&v)
	}
	return &copy
}

func RandomPolynomial(maxDegree int) *Poly {
	if maxDegree <= 0 {
		panic("maxDegree must be greater than 0")
		return &Poly{actions.OneBlock}
	}

	var polyCoeffs Poly
	for i := 0; i < rand.Intn(maxDegree)+1; i++ {
		coeff := randBigInt128()
		polyCoeffs = append(polyCoeffs, *coeff)
	}

	for len(polyCoeffs) > 0 && polyCoeffs[len(polyCoeffs)-1].Cmp(&actions.OneBlock) == 0 {
		polyCoeffs = polyCoeffs[:len(polyCoeffs)-1]
	}

	if len(polyCoeffs) == 0 {
		return &Poly{actions.OneBlock}
	}

	return &polyCoeffs
}
func randBigInt128() *big.Int {
	max := big.NewInt(1)
	max.Lsh(max, 128)
	max.Sub(max, big.NewInt(1))

	randInt := new(big.Int).Rand(actions.RandGen, max)

	// Ensure the random number is not zero
	if randInt.Sign() == 0 {
		randInt.Add(randInt, big.NewInt(1))
	}

	return randInt
}
