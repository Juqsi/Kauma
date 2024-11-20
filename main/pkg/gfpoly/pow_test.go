package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyPowExpected struct {
	Z []string `json:"Z"`
}

func TestGfpolyPow_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyPow
		expected GfpolyPowExpected
	}{
		{
			title: "Basic Test",
			input: GfpolyPow{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				K: 3,
			},
			expected: GfpolyPowExpected{
				Z: []string{
					"AkkAAAAAAAAAAAAAAAAAAA==",
					"DDAAAAAAAAAAAAAAAAAAAA==",
					"LQIIAAAAAAAAAAAAAAAAAA==",
					"8AAAAAAAAAAAAAAAAAAAAA==",
					"ACgCQAAAAAAAAAAAAAAAAA==",
					"AAAMAAAAAAAAAAAAAAAAAA==",
					"AAAAAgAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "exponent 0",
			input: GfpolyPow{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				K: 0,
			},
			expected: GfpolyPowExpected{
				Z: []string{
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Z, testcase.input.Z, "z: Expected %v, got %v", testcase.expected.Z, testcase.input.Z)
		})
	}
}
