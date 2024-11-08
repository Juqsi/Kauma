package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlock2Poly_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Block2Poly
		expected []uint
	}{
		{
			title:    "Basic Test",
			input:    Block2Poly{Semantic: "xex", Block: "ARIAAAAAAAAAAAAAAAAAgA=="},
			expected: []uint{12, 127, 0, 9},
		},
		{
			title:    "Empty Input",
			input:    Block2Poly{Semantic: "xex", Block: ""},
			expected: []uint{},
		}, {
			title:    "Zero Input",
			input:    Block2Poly{Semantic: "xex", Block: "AAAAAAAAAAAAAAAAAAAAAA=="},
			expected: []uint{},
		},
		{
			title:    "Zero Input",
			input:    Block2Poly{Semantic: "gcm", Block: "AAAAAAAAAAAAAAAAAAAAAA=="},
			expected: []uint{},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			response := testcase.input.Result
			assert.ElementsMatch(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
