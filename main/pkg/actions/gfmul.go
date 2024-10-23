package actions

import (
	"Abgabe/main/pkg/utils"
	"math/big"
)

func GfmulBigInt(factor1, factor2, reduce *big.Int) *big.Int {
	result := big.NewInt(0)
	tmpFactor1 := new(big.Int).Set(factor1) // Kopiere factor1, um es zu verändern
	tmpFactor2 := new(big.Int).Set(factor2) // Kopiere factor2, um es zu verändern
	tmpReduce := new(big.Int).Set(reduce)   // Kopiere reduce für die Reduktion

	// Iteriere durch die Bits von factor2
	for tmpFactor2.BitLen() > 0 {
		// Wenn das niedrigste Bit in factor2 gesetzt ist
		if tmpFactor2.Bit(0) == 1 {
			result.Xor(result, tmpFactor1) // XOR mit dem ersten Faktor
		}

		// Shift factor1 nach links (entspricht einer Multiplikation mit x)
		tmpFactor1.Lsh(tmpFactor1, 1)

		// Wenn factor1 größer als das Reduktionspolynom wird, wird reduziert
		if tmpFactor1.BitLen() >= reduce.BitLen() {
			tmpFactor1.Xor(tmpFactor1, tmpReduce) // Polynommodulo-Operation
		}

		// Shift factor2 nach rechts (wir verarbeiten das nächste Bit)
		tmpFactor2.Rsh(tmpFactor2, 1)
	}
	return result
}

func Gfmul(faktor1, faktor2 string) string {
	result := GfmulBigInt(utils.NewLongFromBase64(faktor1).BigInt(), utils.NewLongFromBase64(faktor2).BigInt(), Coeff2Number([]uint{128, 7, 2, 1, 0}))
	return utils.NewLongFromBigInt(result).GetBase64(16)
}
