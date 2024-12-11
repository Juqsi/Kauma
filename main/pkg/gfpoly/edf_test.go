package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyEdfExpected struct {
	Factors [][]string
}

func TestGfpolyEdf_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyEdf
		expected GfpolyEdfExpected
	}{
		{
			title: "Basic Test Task",
			input: GfpolyEdf{
				F: []string{
					"mmAAAAAAAAAAAAAAAAAAAA==",
					"AbAAAAAAAAAAAAAAAAAAAA==",
					"zgAAAAAAAAAAAAAAAAAAAA==",
					"FwAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
				D: 3,
			},
			expected: GfpolyEdfExpected{
				Factors: [][]string{
					{
						"iwAAAAAAAAAAAAAAAAAAAA==",
						"CAAAAAAAAAAAAAAAAAAAAA==",
						"AAAAAAAAAAAAAAAAAAAAAA==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
					},
					{
						"kAAAAAAAAAAAAAAAAAAAAA==",
						"CAAAAAAAAAAAAAAAAAAAAA==",
						"wAAAAAAAAAAAAAAAAAAAAA==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Factors, testcase.input.Factors, "Expected \n %v\n, got\n %v", testcase.expected.Factors, testcase.input.Factors)
		})
	}
}
