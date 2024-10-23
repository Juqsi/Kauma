package testcases

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/models"
	"Abgabe/main/pkg/utils"
	"encoding/json"
	"fmt"
	"os"
)

func GetTestcases() string {
	argsWithProg := os.Args
	if len(argsWithProg) != 2 {
		panic("Wrong number of arguments")
	}
	fmt.Println(argsWithProg)

	jsonData := utils.ReadFile(argsWithProg[1])

	result, err := getTestCasesFromJSON(jsonData)
	if err != nil {
		fmt.Println("Fehler beim Unmarshaling:", err)
		return ""
	}
	return result
}

func getTestCasesFromJSON(jsonData string) (string, error) {
	var testCases models.TestcaseFile
	err := json.Unmarshal([]byte(jsonData), &testCases)
	if err != nil {
		return "", err
	}

	result := make(map[string]interface{})

	for key, testCase := range testCases.Testcases {
		switch testCase.Action {
		case "poly2block":
			var args = new(actions.Poly2Block)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				fmt.Println("Error unmarshalling poly2block arguments:", err)
				continue
			}
			args.Execute()
			result[key] = args

		case "block2poly":
			var args = new(actions.Block2Poly)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				fmt.Println("Error unmarshalling block2poly arguments:", err)
				continue
			}
			args.Execute()
			fmt.Println(args)
			result[key] = args

		case "gfmul":
			var args = new(actions.Gfmul)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				fmt.Println("Error unmarshalling gfmul arguments:", err)
				continue
			}
			args.Execute()
			result[key] = args
		}

	}
	a, _ := json.Marshal(result)
	fmt.Println(string(a))
	return string(a), nil
}
