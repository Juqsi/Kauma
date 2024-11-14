package testcases

import (
	"Abgabe/main/pkg/models"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
)

func runTestcases(testCases models.TestcaseFile) (string, error) {
	result := make(map[string]map[string]interface{})

	handlers := map[string]func([]byte) (map[string]interface{}, error){
		"poly2block":     handlePoly2Block,
		"block2poly":     handleBlock2Poly,
		"gfmul":          handleGfmul,
		"sea128":         handleSea128,
		"xex":            handleXex,
		"padding_oracle": handlePaddingOracle,
		"gcm_encrypt":    handleGcmEncrypt,
		"gcm_decrypt":    handleGcmDecrypt,
		"gfpoly_add":     handleGfpolyAdd,
		"gfpoly_mul":     handleGfpolyMul,
		"gfpoly_pow":     handleGfpolyPow,
		"gfdiv":          handleGfdiv,
	}

	for key, testCase := range testCases.Testcases {
		func(key string, testCase models.Testcase) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Fprintf(os.Stderr, "Error in testcase %s: %v\n", testCase.Arguments, r)
					fmt.Fprintf(os.Stderr, "Stacktrace:\n%s\n", debug.Stack())
				}
			}()

			if handler, found := handlers[testCase.Action]; found {
				if res, err := handler(testCase.Arguments); err == nil {
					result[key] = res
				} else {
					fmt.Fprintf(os.Stderr, "(Marshal-) Error in testcase %s: %v\n", key, err)
				}
			} else {
				fmt.Fprintf(os.Stderr, "Unknown action: %s\n", testCase.Action)
			}
		}(key, testCase)
	}

	res := struct {
		Response map[string]map[string]interface{} `json:"responses"`
	}{Response: result}

	a, _ := json.Marshal(res)
	return string(a), nil
}
