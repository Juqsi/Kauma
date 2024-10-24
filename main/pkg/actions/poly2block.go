package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Poly2Block struct {
	Semantic     string `json:"Semantic"`
	Coefficients []uint `json:"Coefficients"`
	Result       string `json:"Block"`
}

func (p *Poly2Block) Execute() {
	p.Result = utils.NewLongFromBigInt(Coeff2Number(p.Coefficients)).GetBase64(-1)
}

func Coeff2Number(coefficients []uint) *big.Int {
	number := new(big.Int)
	for _, coeff := range coefficients {
		number.SetBit(number, int(coeff), 1)
	}
	return number
}
