package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyDiffExpected struct {
	FStrich []string `json:"Z"`
}

func TestGfpolyDiff_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyDiff
		expected GfpolyDiffExpected
	}{
		{
			title: "Basic Test Encode ",
			input: GfpolyDiff{
				F: []string{
					"IJustWannaTellYouAAAAA==",
					"HowImFeelingAAAAAAAAAA==",
					"GottaMakeYouAAAAAAAAAA==",
					"UnderstaaaaaaaaaaaaanQ==",
				},
			},
			expected: GfpolyDiffExpected{
				FStrich: []string{
					"HowImFeelingAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"UnderstaaaaaaaaaaaaanQ==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.FStrich, testcase.input.FStrich, "Expected \n %v\n, got\n %v", testcase.expected.FStrich, testcase.input.FStrich)
		})
	}
}
