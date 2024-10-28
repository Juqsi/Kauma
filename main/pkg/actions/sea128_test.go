package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSea128_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    sea128
		expected string
	}{
		{
			title:    "Basic Test Encode",
			input:    sea128{Mode: "encrypt", Key: "istDASeincoolerKEYrofg==", Input: "yv66vvrO263eyviIiDNEVQ=="},
			expected: "D5FDo3iVBoBN9gVi9/MSKQ==",
		}, {
			title:    "Basic Test Decode",
			input:    sea128{Mode: "decrypt", Key: "istDASeincoolerKEYrofg==", Input: "D5FDo3iVBoBN9gVi9/MSKQ=="},
			expected: "yv66vvrO263eyviIiDNEVQ==",
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
