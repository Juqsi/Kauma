package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

/*func Poly2Block(coefficients []uint) string {
	number := Coeff2Number(coefficients)
	return utils.NewLongFromBigInt(number).GetBase64(-1)
}
*/

type Poly2Block struct {
	Semantic     string `json:"semantic"`
	Coefficients []uint `json:"coefficients"`
	Result       string `json:"block"`
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
