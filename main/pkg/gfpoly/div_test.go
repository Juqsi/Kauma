package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyDivExpected struct {
	R []string `json:"R"`
	Q []string `json:"Q"`
}

func TestGfpolyDiv_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyDiv
		expected GfpolyDivExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyDiv{
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
			expected: GfpolyDivExpected{
				Q: []string{
					"nAIAgCAIAgCAIAgCAIAgCg==",
					"m85znOc5znOc5znOc5znOQ==",
				},
				R: []string{
					"lQNA0DQNA0DQNA0DQNA0Dg==",
				},
			},
		},
		{
			title: "Zero divided by something",
			input: GfpolyDiv{
				A: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "Zero divided by something",
			input: GfpolyDiv{
				A: []string{
					"IAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"IAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Q, testcase.input.Q, "Q: Expected %v, got %v", testcase.expected.Q, testcase.input.Q)
			assert.Equal(t, testcase.expected.R, testcase.input.R, "R: Expected %v, got %v", testcase.expected.R, testcase.input.R)
		})
	}
}
