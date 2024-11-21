package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolymakeMonicExpected struct {
	ASternchen []string
}

func TestGfpolymakeMonic_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyMakeMonic
		expected GfpolymakeMonicExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyMakeMonic{
				A: []string{
					"NeverGonnaGiveYouUpAAA==",
					"NeverGonnaLetYouDownAA==",
					"NeverGonnaRunAroundAAA==",
					"AndDesertYouAAAAAAAAAA==",
				},
			},
			expected: GfpolymakeMonicExpected{
				ASternchen: []string{
					"edY47onJ4MtCENDTHG/sZw==",
					"oaXjCKnceBIxSavZ9eFT8w==",
					"1Ial5rAJGOucIdUe3zh5bw==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.ASternchen, testcase.input.ASternchen, "Expected \n %v\n, got\n %v", testcase.expected.ASternchen, testcase.input.ASternchen)
		})
	}
}
