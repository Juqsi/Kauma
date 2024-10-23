package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlock2Poly(t *testing.T) {
	testcases := []struct {
		title    string
		input    string
		expected []uint
	}{
		{
			title:    "Basic Test",
			input:    "ARIAAAAAAAAAAAAAAAAAgA==",
			expected: []uint{12, 127, 0, 9},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			response := Block2poly(testcase.input)
			assert.ElementsMatch(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
