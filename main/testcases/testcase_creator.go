package testcases

import (
	"Abgabe/main/pkg/models"
	"Abgabe/main/pkg/utils"
	"encoding/json"
	"os"
)

func RunTestcases() (string, bool) {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		panic("Wrong number of arguments")
	}
	jsonData := utils.ReadFile(argsWithProg[1])

	testcases, err := getTestCasesFromJSON(jsonData)
	if err != nil {
		panic("Fehler beim Unmarshaling: " + err.Error())
	}
	return runTestcases(testcases)
}

func getTestCasesFromJSON(jsonData string) (models.TestcaseFile, error) {
	var testCases models.TestcaseFile
	err := json.Unmarshal([]byte(jsonData), &testCases)
	if err != nil {
		return models.TestcaseFile{}, err
	}
	return testCases, nil
}
