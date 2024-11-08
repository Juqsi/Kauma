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
	switch p.Semantic {
	case "xex":
		p.Result = utils.NewLongFromBigInt(Coeff2Number(p.Coefficients)).GetLittleEndianInBase64(16)
	case "gcm":
		p.Result = utils.NewLongFromBigInt(Coeff2Number(p.Coefficients)).Reverse(128).GetBase64(16)
	}

}

func Coeff2Number(coefficients []uint) big.Int {
	number := new(big.Int)
	for _, coeff := range coefficients {
		number.SetBit(number, int(coeff), 1)
	}
	return *number
}
