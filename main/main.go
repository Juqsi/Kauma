package main

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/utils"
	"fmt"
)

func main() {
	//fmt.Println(testcases.GetTestcases())

	Coefficients := []uint{12, 127, 0, 9}
	Poly := actions.Poly2Block(Coefficients)
	fmt.Println(utils.NewLongFromBigInt(Poly).GetBigEndianInBase64())
	fmt.Println("ARIAAAAAAAAAAAAAAAAAgA==")

	PolyString := "ARIAAAAAAAAAAAAAAAAAgA=="
	fmt.Println(Poly.Text(2))
	Coefficients = actions.Number2Coefficients(utils.NewLongFromBase64InBigEndian(PolyString).BigInt())
	fmt.Println(Coefficients)
	fmt.Println("[12, 127, 0, 9]")

}
