package main

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/testcases"
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(int64(0x49))
	b := new(big.Int).SetInt64(int64(0xb1))

	fmt.Println(actions.GfmulBigInt(a, b, actions.Coeff2Number([]uint{8, 6, 5, 2, 0})).Text(2))
	tests := testcases.RunTestcases()
	fmt.Println(tests)
}
