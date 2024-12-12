package testcases

import (
	"Abgabe/main/pkg/models"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
)

func runTestcases(testCases models.TestcaseFile) string {
	result := make(map[string]map[string]interface{}, len(testCases.Testcases))
	var mu sync.Mutex
	var wg sync.WaitGroup

	// Handler-Funktionen
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
		"gfpoly_powmod":     handleGfpolyPowMod,
		"gfpoly_sort":       handleGfpolySort,
		"gfpoly_make_monic": handleGfpolyMakeMonic,
		"gfpoly_sqrt":       handleGfpolySqrt,
		"gfpoly_diff":       handleGfpolyDiff,
		"gfpoly_gcd":        handleGfpolyGgt,
		"gfpoly_factor_sff": handleGfpolySff,
		"gfpoly_factor_ddf": handleGfpolyDdf,
		"gfpoly_factor_edf": handleGfpolyEdf,
	}

	type Job struct {
		Key      string
		TestCase models.Testcase
	}

	jobQueue := make(chan Job, len(testCases.Testcases))

	numCPUs := runtime.NumCPU() * 2
	numWorkers := numCPUs

	worker := func() {
		defer wg.Done()
		for job := range jobQueue {
			key := job.Key
			testCase := job.TestCase

			defer func() {
				if r := recover(); r != nil {
					mu.Lock()
					_, _ = fmt.Fprintf(os.Stderr, "Error in testcase \n action: %s \n Arguments: %s: %v\n", testCase.Action, testCase.Arguments, r)
					_, _ = fmt.Fprintf(os.Stderr, "Stacktrace:\n%s\n", debug.Stack())
					mu.Unlock()
				}
			}()

			if handler, found := handlers[testCase.Action]; found {

				if res, err := handler(testCase.Arguments); err == nil {
					mu.Lock()
					result[key] = res
					mu.Unlock()
				} else {
					_, _ = fmt.Fprintf(os.Stderr, "(Marshal-) Error in testcase %s: %v\n", key, err)
				}
			} else {
				_, _ = fmt.Fprintf(os.Stderr, "Unknown action: %s\n", testCase.Action)
			}
		}
	}

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	for key, testCase := range testCases.Testcases {
		jobQueue <- Job{Key: key, TestCase: testCase}
	}
	close(jobQueue)

	wg.Wait()

	res := struct {
		Response map[string]map[string]interface{} `json:"responses"`
	}{Response: result}

	a, _ := json.Marshal(res)

	return string(a)
}
