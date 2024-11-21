package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSea128_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Sea128
		expected string
	}{
		{
			title:    "Basic Test Encode",
			input:    Sea128{Mode: "encrypt", Key: "istDASeincoolerKEYrofg==", Input: "yv66vvrO263eyviIiDNEVQ=="},
			expected: "D5FDo3iVBoBN9gVi9/MSKQ==",
		}, {
			title:    "Basic Test Decode",
			input:    Sea128{Mode: "decrypt", Key: "istDASeincoolerKEYrofg==", Input: "D5FDo3iVBoBN9gVi9/MSKQ=="},
			expected: "yv66vvrO263eyviIiDNEVQ==",
		},
		{
			title:    "Empty Input Encode",
			input:    Sea128{Mode: "encrypt", Key: "istDASeincoolerKEYrofg==", Input: ""},
			expected: "/35bimzmC3HOC0HFLpjkCw==",
		}, {
			title:    "Empty Input Decode",
			input:    Sea128{Mode: "decrypt", Key: "istDASeincoolerKEYrofg==", Input: ""},
			expected: "rCGMWXjfwGVQU4cx1C2npw==",
		}, {
			title:    "Failed Testcase",
			input:    Sea128{Mode: "decrypt", Key: "vyG55UYdcvvBSQ75h9R2HQ==", Input: "h4nTU4jnW2SW+kt74Pc0Ow=="},
			expected: "AJDIJykK8cy3JDbSf+8yFA==",
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
