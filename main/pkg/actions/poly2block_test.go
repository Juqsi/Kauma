package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoly2Block_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Poly2Block
		expected string
	}{
		{
			title:    "Basic Test",
			input:    Poly2Block{Coefficients: []uint{12, 127, 0, 9}},
			expected: "ARIAAAAAAAAAAAAAAAAAgA==",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			block := testcase.input.Result
			assert.Equal(t, testcase.expected, block, "Expected %v, got %v", testcase.expected, block)
		})
	}
}
