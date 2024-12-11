package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type GfpolyPowModExpected struct {
	Z []string `json:"Z"`
}

func TestGfpolyPowMod_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyPowMod
		expected GfpolyPowModExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyPowMod{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				M: []string{
					"KryptoanalyseAAAAAAAAA==",
					"DHBWMannheimAAAAAAAAAA==",
				},
				K: *big.NewInt(1000),
			},
			expected: GfpolyPowModExpected{
				Z: []string{
					"oNXl5P8xq2WpUTP92u25zg==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Z, testcase.input.Z, "z: Expected %v, got %v", testcase.expected.Z, testcase.input.Z)
		})
	}
}
