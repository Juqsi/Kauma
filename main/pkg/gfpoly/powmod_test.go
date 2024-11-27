package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type GfpolyPowmodExpected struct {
	Z []string `json:"Z"`
}

func TestGfpolyPowmod_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyPowmod
		expected GfpolyPowmodExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyPowmod{
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
			expected: GfpolyPowmodExpected{
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
