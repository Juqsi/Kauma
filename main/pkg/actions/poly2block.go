package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Polynom struct {
	Semantic     string `json:"semantic"`
	Coefficients []uint `json:"coefficients"`
	Block        string `json:"block"`
	Number       *big.Int
}

func Poly2Block(coefficients []uint) string {
	number := Coeff2Number(coefficients)
	return utils.NewLongFromBigInt(number).GetBigEndianInBase64()
}

func Coeff2Number(coefficients []uint) *big.Int {
	number := new(big.Int)
	for _, coeff := range coefficients {
		tmp := new(big.Int).Lsh(big.NewInt(1), coeff)
		number = new(big.Int).Or(number, tmp)
	}
	return number
}
