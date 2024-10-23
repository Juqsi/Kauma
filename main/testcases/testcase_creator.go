package testcases

import (
	"Abgabe/main/pkg/models"
	utils2 "Abgabe/main/pkg/utils"
	"fmt"
	"os"
)

func GetTestcases() []models.Testcase {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		panic("Wrong number of arguments")
	}

	json := utils2.ReadFile(argsWithProg[1])

	result, err := utils2.GetTestCasesFromJSON(json)
	if err != nil {
		fmt.Println("Fehler beim Unmarshaling:", err)
		return nil
	}
	return result
}
