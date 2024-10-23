package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGfmul(t *testing.T) {
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
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			response := testcase.input.Result
			assert.Equal(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
