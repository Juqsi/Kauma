package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

type Block2Poly struct {
	Semantic string `json:"Semantic"`
	Block    string `json:"block"`
	Result   []uint `json:"Coefficients"`
}

func (b *Block2Poly) Execute() {
	switch b.Semantic {
	case "xex":
		b.Result = XexNumber2Coefficients(utils.NewLongFromLittleEndianInBase64(b.Block).Int)
	case "gcm":
		b.Result = XexNumber2Coefficients(utils.NewBigEndianLongFromGcmInBase64(b.Block).Int)
	}
}

func XexNumber2Coefficients(number big.Int) []uint {
	ret := []uint{}
	bitLen := number.BitLen()
	for i := 0; i < bitLen; i++ {
		if number.Bit(i) == 1 {
			ret = append(ret, uint(i))
		}
	}
	return ret
}
