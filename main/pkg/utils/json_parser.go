package utils

/*
func GetTestCasesFromJSON(jsonData string) ([]models.Testcase, error) {
	var testCasesMap models.TestcaseFile

	err := json.Unmarshal([]byte(jsonData), &testCasesMap)
	if err != nil {
		return nil, err
	}
	var testCases []models.Testcase
	for key, testCase := range testCasesMap.Testcases {
		testCase.Key = key

		argType, _ := getArgumentStructAndTest(testCase.Action)

		//testCase.Execute = test;

		argBytes, err := json.Marshal(testCase.Arguments) // Arguments in JSON umwandeln
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(argBytes, &argType)
		if err != nil {
			return nil, err
		}

		testCase.Arguments = argType

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}

func getArgumentStructAndTest(action string) (interface{}, func(interface{}) (interface{}, error)) {
	switch action {
	case "poly2block":
		return actions.Polynom{}, nil //actions.Poly2block
	case "block2poly":
		return actions.Polynom{}, nil //actions.Block2poly
	default:
		panic("Unbekannte Aktion: " + action)
	}
}

*/
