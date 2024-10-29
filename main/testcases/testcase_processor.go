package testcases

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/models"
	"encoding/json"
)

func runTestcases(testCases models.TestcaseFile) (string, error) {
	result := make(map[string]map[string]interface{})

	for key, testCase := range testCases.Testcases {
		switch testCase.Action {
		case "poly2block":
			var args = new(actions.Poly2Block)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				continue
			}
			args.Execute()
			result[key] = map[string]interface{}{"block": args.Result}

		case "block2poly":
			var args = new(actions.Block2Poly)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				continue
			}
			args.Execute()
			result[key] = map[string]interface{}{"coefficients": args.Result}

		case "gfmul":
			var args = new(actions.Gfmul)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				continue
			}
			args.Execute()
			result[key] = map[string]interface{}{"product": args.Result}
		case "sea128":
			var args = new(actions.Sea128)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				continue
			}
			args.Execute()
			result[key] = map[string]interface{}{"output": args.Result}
		case "xex":
			var args = new(actions.Xex)
			if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
				continue
			}
			args.Execute()
			result[key] = map[string]interface{}{"output": args.Result}
		}

	}
	a, _ := json.Marshal(result)
	return string(a), nil
}
