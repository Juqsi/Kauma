package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyDdfExpected struct {
	Factors FactorsModelWithDegree
}

func TestGfpolyDdf_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyDdf
		expected GfpolyDdfExpected
	}{
		{
			title: "Basic Test Task",
			input: GfpolyDdf{
				F: []string{
					"tpkgAAAAAAAAAAAAAAAAAA==",
					"m6MQAAAAAAAAAAAAAAAAAA==",
					"8roAAAAAAAAAAAAAAAAAAA==",
					"3dUAAAAAAAAAAAAAAAAAAA==",
					"FwAAAAAAAAAAAAAAAAAAAA==",
					"/kAAAAAAAAAAAAAAAAAAAA==",
					"a4AAAAAAAAAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDdfExpected{
				Factors: FactorsModelWithDegree{
					{
						Factor: []string{
							"q4AAAAAAAAAAAAAAAAAAAA==",
							"gAAAAAAAAAAAAAAAAAAAAA==",
						},
						Degree: 1,
					},
					{
						Factor: []string{
							"mmAAAAAAAAAAAAAAAAAAAA==",
							"AbAAAAAAAAAAAAAAAAAAAA==",
							"zgAAAAAAAAAAAAAAAAAAAA==",
							"FwAAAAAAAAAAAAAAAAAAAA==",
							"AAAAAAAAAAAAAAAAAAAAAA==",
							"wAAAAAAAAAAAAAAAAAAAAA==",
							"gAAAAAAAAAAAAAAAAAAAAA==",
						},
						Degree: 3,
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Factors, testcase.input.Factors, "Expected \n %v\n, got\n %v", testcase.expected.Factors, testcase.input.Factors)
		})
	}
}
