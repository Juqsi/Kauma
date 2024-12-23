package glasskey

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type prngExpected struct {
	blocks []string
}

func TestPrng_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Prng
		expected prngExpected
	}{
		{
			title: "Basic Test Task",
			input: Prng{
				AgencyKey: "T01HV1RG",
				Seed:      "ur1EoxDElJs=",
				Lengths: []int{
					4,
					8,
					13,
					12,
					1,
					9,
				},
			},
			expected: prngExpected{
				blocks: []string{
					"9q32ZQ==",
					"r2I4mx+M13E=",
					"RvMXtSbjaKkuBXoUsQ==",
					"gwdanlsoBDPlMuzk",
					"ZQ==",
					"P3ixhiNIbxur",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.blocks, testcase.input.Blocks, "blocks: Expected \n %v\n, got\n %v", testcase.expected.blocks, testcase.input.Blocks)
		})
	}
}
