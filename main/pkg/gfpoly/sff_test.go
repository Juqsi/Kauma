package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolySffExpected struct {
	Factors FactorsModel
}

func TestGfpolySff_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolySff
		expected GfpolySffExpected
	}{
		{
			title: "Basic Test Task",
			input: GfpolySff{
				F: []string{
					"vL77UwAAAAAAAAAAAAAAAA==",
					"mEHchYAAAAAAAAAAAAAAAA==",
					"9WJa0MAAAAAAAAAAAAAAAA==",
					"akHfwWAAAAAAAAAAAAAAAA==",
					"E12o/QAAAAAAAAAAAAAAAA==",
					"vKJ/FgAAAAAAAAAAAAAAAA==",
					"yctWwAAAAAAAAAAAAAAAAA==",
					"c1BXYAAAAAAAAAAAAAAAAA==",
					"o0AtAAAAAAAAAAAAAAAAAA==",
					"AbP2AAAAAAAAAAAAAAAAAA==",
					"k2YAAAAAAAAAAAAAAAAAAA==",
					"vBYAAAAAAAAAAAAAAAAAAA==",
					"dSAAAAAAAAAAAAAAAAAAAA==",
					"69gAAAAAAAAAAAAAAAAAAA==",
					"VkAAAAAAAAAAAAAAAAAAAA==",
					"a4AAAAAAAAAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolySffExpected{
				Factors: FactorsModel{
					{
						Factor: []string{
							"q4AAAAAAAAAAAAAAAAAAAA==",
							"gAAAAAAAAAAAAAAAAAAAAA==",
						},
						Exponent: 1,
					},
					{
						Factor: []string{
							"iwAAAAAAAAAAAAAAAAAAAA==",
							"CAAAAAAAAAAAAAAAAAAAAA==",
							"AAAAAAAAAAAAAAAAAAAAAA==",
							"gAAAAAAAAAAAAAAAAAAAAA==",
						},
						Exponent: 2,
					},
					{
						Factor: []string{
							"kAAAAAAAAAAAAAAAAAAAAA==",
							"CAAAAAAAAAAAAAAAAAAAAA==",
							"wAAAAAAAAAAAAAAAAAAAAA==",
							"gAAAAAAAAAAAAAAAAAAAAA==",
						},
						Exponent: 3,
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
