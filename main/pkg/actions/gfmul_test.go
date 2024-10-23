package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGfmul(t *testing.T) {
	testcases := []struct {
		title    string
		input    [2]string
		expected string
	}{
		{
			title:    "Basic Test",
			input:    [2]string{"ARIAAAAAAAAAAAAAAAAAgA==", "AgAAAAAAAAAAAAAAAAAAAA=="},
			expected: "hSQAAAAAAAAAAAAAAAAAAA==",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			response := Gfmul(testcase.input[0], testcase.input[1])
			assert.Equal(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
