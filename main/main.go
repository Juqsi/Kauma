package main

import (
	"Abgabe/main/testcases"
	"fmt"
	"os"
)

func main() {
	tests, err := testcases.RunTestcases()
	fmt.Println(tests)
	if err {
		os.Exit(2)
	}
}
