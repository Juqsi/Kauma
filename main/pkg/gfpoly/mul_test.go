package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyMulExpected struct {
	P []string `json:"Z"`
}

func TestGfpolyMul_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyMul
		expected GfpolyMulExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyMul{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyMulExpected{
				P: []string{
					"MoAAAAAAAAAAAAAAAAAAAA==",
					"sUgAAAAAAAAAAAAAAAAAAA==",
					"MbQAAAAAAAAAAAAAAAAAAA==",
					"AAhAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "Lukas debug",
			input: GfpolyMul{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyMulExpected{
				P: []string{
					"CCAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"oAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAIAAAAAAAAAAAAAAAAAA=="},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.P, testcase.input.P, "P: Expected \n %v\n, got\n %v", testcase.expected.P, testcase.input.P)
		})
	}
}
