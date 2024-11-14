package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGfdiv_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Gfdiv
		expected string
	}{
		{
			title:    "Basic Test",
			input:    Gfdiv{Factor1: "JAAAAAAAAAAAAAAAAAAAAA==", Factor2: "wAAAAAAAAAAAAAAAAAAAAA=="},
			expected: "OAAAAAAAAAAAAAAAAAAAAA==",
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
