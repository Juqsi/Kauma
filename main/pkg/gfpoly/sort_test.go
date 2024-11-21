package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolySortExpected struct {
	SortedPolys [][]string
}

func TestGfpolySort_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolySort
		expected GfpolySortExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolySort{
				Polys: [][]string{
					{
						"NeverGonnaGiveYouUpAAA==",
						"NeverGonnaLetYouDownAA==",
						"NeverGonnaRunAroundAAA==",
						"AndDesertYouAAAAAAAAAA==",
					},
					{
						"WereNoStrangersToLoveA==",
						"YouKnowTheRulesAAAAAAA==",
						"AndSoDoIAAAAAAAAAAAAAA==",
					},
					{
						"NeverGonnaMakeYouCryAA==",
						"NeverGonnaSayGoodbyeAA==",
						"NeverGonnaTellALieAAAA==",
						"AndHurtYouAAAAAAAAAAAA==",
					},
				},
			},
			expected: GfpolySortExpected{
				SortedPolys: [][]string{
					{
						"WereNoStrangersToLoveA==",
						"YouKnowTheRulesAAAAAAA==",
						"AndSoDoIAAAAAAAAAAAAAA==",
					},
					{
						"NeverGonnaMakeYouCryAA==",
						"NeverGonnaSayGoodbyeAA==",
						"NeverGonnaTellALieAAAA==",
						"AndHurtYouAAAAAAAAAAAA==",
					},
					{
						"NeverGonnaGiveYouUpAAA==",
						"NeverGonnaLetYouDownAA==",
						"NeverGonnaRunAroundAAA==",
						"AndDesertYouAAAAAAAAAA==",
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.SortedPolys, testcase.input.SortedPolys, "Expected %v, got %v", testcase.expected.SortedPolys, testcase.input.SortedPolys)
		})
	}
}
