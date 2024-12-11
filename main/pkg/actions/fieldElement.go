package actions

import (
	"math/big"
	"math/rand"
	"time"
)

var (
	RandGen   *rand.Rand
	Reduce128 big.Int
	OneBlock  big.Int
)

func init() {
	RandGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	Reduce128 = Coeff2Number([]uint{128, 7, 2, 1, 0})
	OneBlock = *big.NewInt(1)
}
