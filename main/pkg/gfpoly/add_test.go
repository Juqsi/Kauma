package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyAddExpected struct {
	S []string `json:"Z"`
}

func TestGfpolyAdd_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyAdd
		expected GfpolyAddExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyAdd{
				A: []string{
					"NeverGonnaGiveYouUpAAA==",
					"NeverGonnaLetYouDownAA==",
					"NeverGonnaRunAroundAAA==",
					"AndDesertYouAAAAAAAAAA==",
				},
				B: []string{
					"KryptoanalyseAAAAAAAAA==",
					"DHBWMannheimAAAAAAAAAA==",
				},
			},
			expected: GfpolyAddExpected{
				S: []string{
					"H1d3GuyA9/0OxeYouUpAAA==",
					"OZuIncPAGEp4tYouDownAA==",
					"NeverGonnaRunAroundAAA==",
					"AndDesertYouAAAAAAAAAA==",
				},
			},
		},
		{
			title: "two empty polynomials",
			input: GfpolyAdd{
				A: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyAddExpected{
				S: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "two empty polynomials",
			input: GfpolyAdd{
				A: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==", "AAAAAAAAAAAAAAAAAAAAAA==", "AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyAddExpected{
				S: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.S, testcase.input.Z, "Z: Expected %v, got %v", testcase.expected.S, testcase.input.Z)
		})
	}
}
