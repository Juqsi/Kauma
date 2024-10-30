package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGfmul_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Gfmul
		expected string
	}{
		{
			title:    "Basic Test",
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: "AgAAAAAAAAAAAAAAAAAAAA==", Semantic: "xex"},
			expected: "hSQAAAAAAAAAAAAAAAAAAA==",
		},
		{
			title:    "Empty Input",
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: "", Semantic: "xex"},
			expected: "AAAAAAAAAAAAAAAAAAAAAA==",
		},
		{
			title:    "Zero Multiplication",
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: "AAAAAAAAAAAAAAAAAAAAAA==", Semantic: "xex"},
			expected: "AAAAAAAAAAAAAAAAAAAAAA==",
		},
		{
			title:    "All 1",
			input:    Gfmul{Factor2: "/////////////////////w==", Factor1: "/////////////////////w==", Semantic: "xex"},
			expected: "L0BVVVVVVVVVVVVVVVVVVQ==",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			response := testcase.input.Result
			assert.Equal(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
