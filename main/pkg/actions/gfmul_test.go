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
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: "AgAAAAAAAAAAAAAAAAAAAA=="},
			expected: "hSQAAAAAAAAAAAAAAAAAAA==",
		},
		{
			title:    "Empty Input",
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: ""},
			expected: "AAAAAAAAAAAAAAAAAAAAAA==",
		},
		{
			title:    "Empty Input",
			input:    Gfmul{Factor2: "ARIAAAAAAAAAAAAAAAAAgA==", Factor1: "AAAAAAAAAAAAAAAAAAAAAA=="},
			expected: "AAAAAAAAAAAAAAAAAAAAAA==",
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
