package testcases

import (
	"Abgabe/main/pkg/models"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
)

func runTestcases(testCases models.TestcaseFile) (string, bool) {
	result := make(map[string]map[string]interface{})
	handlerCounts := make(map[string]int)
	errorOccured := false

	handlers := map[string]func([]byte) (map[string]interface{}, error){
		"poly2block":        handlePoly2Block,
		"block2poly":        handleBlock2Poly,
		"gfmul":             handleGfmul,
		"sea128":            handleSea128,
		"xex":               handleXex,
		"padding_oracle":    handlePaddingOracle,
		"gcm_encrypt":       handleGcmEncrypt,
		"gcm_decrypt":       handleGcmDecrypt,
		"gfpoly_add":        handleGfpolyAdd,
		"gfpoly_mul":        handleGfpolyMul,
		"gfpoly_pow":        handleGfpolyPow,
		"gfdiv":             handleGfdiv,
		"gfpoly_divmod":     handleGfpolyDiv,
		"gfpoly_powmod":     handleGfpolyPowmod,
		"gfpoly_sort":       handleGfpolySort,
		"gfpoly_make_monic": handleGfpolyMakeMonic,
		"gfpoly_sqrt":       handleGfpolySqrt,
		"gfpoly_diff":       handleGfpolyDiff,
		"gfpoly_gcd":        handleGfpolyGgt,
	}

	for key, testCase := range testCases.Testcases {
		func(key string, testCase models.Testcase) {
			defer func() {
				if r := recover(); r != nil {
					errorOccured = true
					_, _ = fmt.Fprintf(os.Stderr, "Error in testcase \n action: %s \n Arguments: %s: %v\n", testCase.Action, testCase.Arguments, r)
					_, _ = fmt.Fprintf(os.Stderr, "Stacktrace:\n%s\n", debug.Stack())
					handlerCounts[testCase.Action+"-recovered"]++
				}
			}()

			if handler, found := handlers[testCase.Action]; found {
				handlerCounts[testCase.Action]++
				if res, err := handler(testCase.Arguments); err == nil {
					result[key] = res
				} else {
					_, _ = fmt.Fprintf(os.Stderr, "(Marshal-) Error in testcase %s: %v\n", key, err)
				}
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "Unknown action: %s\n", testCase.Action)
			}
		}(key, testCase)
	}

	res := struct {
		Response map[string]map[string]interface{} `json:"responses"`
	}{Response: result}

	a, _ := json.Marshal(res)

	if errorOccured {
		stats, _ := json.Marshal(handlerCounts)
		_, _ = fmt.Fprintf(os.Stderr, "Statistik: \n %s", stats)
	}

	return string(a), errorOccured
}
