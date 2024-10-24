package main

import (
	"Abgabe/main/testcases"
	"fmt"
)

func main() {
	tests := testcases.RunTestcases()
	fmt.Println(tests)
}
