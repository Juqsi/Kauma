package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolySqrtExpected struct {
	S []string
}

func TestGfpolySqrt_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolySqrt
		expected GfpolySqrtExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolySqrt{
				Q: []string{
					"5TxUxLHO1lHE/rSFquKIAg==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"0DEUJYdHlmd4X7nzzIdcCA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"PKUa1+JHTxHE8y3LbuKIIA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"Ds96KiAKKoigKoiKiiKAiA==",
				},
			},
			expected: GfpolySqrtExpected{
				S: []string{
					"NeverGonnaGiveYouUpAAA==",
					"NeverGonnaLetYouDownAA==",
					"NeverGonnaRunAroundAAA==",
					"AndDesertYouAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.S, testcase.input.S, "Expected \n %v\n, got\n %v", testcase.expected.S, testcase.input.S)
		})
	}
}
