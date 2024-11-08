package testcases

import (
	"Abgabe/main/pkg/actions"
	"Abgabe/main/pkg/models"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
)

func runTestcases(testCases models.TestcaseFile) (string, error) {

	result := make(map[string]map[string]interface{})

	for key, testCase := range testCases.Testcases {
		func(key string, testCase models.Testcase) {
			defer func() {
				if r := recover(); r != nil {
					_, _ = fmt.Fprintf(os.Stderr, "Error in testcase %s: %v\n", testCase.Arguments, r)
					_, _ = fmt.Fprintf(os.Stderr, "Stacktrace:\n%s\n", debug.Stack())
				}
			}()

			switch testCase.Action {
			case "poly2block":
				var args = new(actions.Poly2Block)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"block": args.Result}

			case "block2poly":
				var args = new(actions.Block2Poly)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"coefficients": args.Result}

			case "gfmul":
				var args = new(actions.Gfmul)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"product": args.Result}

			case "sea128":
				var args = new(actions.Sea128)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"output": args.Result}

			case "xex":
				var args = new(actions.Xex)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"output": args.Result}

			case "padding_oracle":
				var args = new(actions.PaddingOracle)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"plaintext": args.Result}

			case "gcm_encrypt":
				var args = new(actions.Gcm_Encrypt)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"ciphertext": args.Ciphertext,
					"tag": args.Tag,
					"L":   args.L,
					"H":   args.H}
			case "gcm_decrypt":
				var args = new(actions.Gcm_Decrypt)
				if err := json.Unmarshal(testCase.Arguments, &args); err != nil {
					return
				}
				args.Execute()
				result[key] = map[string]interface{}{"authentic": args.Authentic,
					"plaintext": args.Plaintext}
			}

		}(key, testCase)
	}

	res := struct {
		Response map[string]map[string]interface{} `json:"responses"`
	}{Response: result}
	a, _ := json.Marshal(res)
	return string(a), nil
}
