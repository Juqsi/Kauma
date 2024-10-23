package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoly2Block(t *testing.T) {
	testcases := []struct {
		title    string
		input    []uint
		expected string
	}{
		{
			title:    "Basic Test",
			input:    []uint{12, 127, 0, 9},
			expected: "ARIAAAAAAAAAAAAAAAAAgA==",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			block := Poly2Block(testcase.input)
			assert.Equal(t, testcase.expected, block, "Expected %v, got %v", testcase.expected, block)
		})
	}
}
