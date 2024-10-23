package actions

import (
	"Abgabe/main/pkg/utils"
	"fmt"
	"math/big"
)

func GfmulBigInt(faktor1, faktor2, reducePoly *big.Int) *big.Int {
	println("Faktor1: " + faktor1.Text(2))
	fmt.Println(Number2Coefficients(faktor1))
	println("Faktor2: " + faktor2.Text(2))
	fmt.Println(Number2Coefficients(faktor2))
	println("reducer: " + reducePoly.Text(2))
	result := new(big.Int).Xor(faktor1, faktor2)
	if reducePoly.BitLen() <= result.BitLen() {
		result = new(big.Int).Xor(result, reducePoly)
	}
	tmp := new(big.Int).And(faktor1, faktor2)
	tmp = new(big.Int).Lsh(tmp, 1)
	fmt.Println("Tmp vor reduce: " + tmp.Text(2))
	if reducePoly.BitLen() <= tmp.BitLen() {
		tmp = new(big.Int).Xor(tmp, reducePoly)
	}
	fmt.Println("Tmp nac reduce: " + tmp.Text(2))
	result = new(big.Int).Or(result, tmp)
	fmt.Println("Result: " + result.Text(2))
	return result
}

func Gfmul(faktor1, faktor2 string) string {
	return utils.NewLongFromBigInt(GfmulBigInt(utils.NewLongFromBase64InBigEndian(faktor1).BigInt(), utils.NewLongFromBase64InBigEndian(faktor2).BigInt(), Coeff2Number([]uint{128, 7, 2, 1, 0}))).GetBigEndianInBase64()
}
