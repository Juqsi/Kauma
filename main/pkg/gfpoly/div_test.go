package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyDivExpected struct {
	R []string `json:"R"`
	Q []string `json:"Q"`
}

func TestGfpolyDiv_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyDiv
		expected GfpolyDivExpected
	}{
		{
			title: "Basic Test Encode aes",
			input: GfpolyDiv{
				A: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"nAIAgCAIAgCAIAgCAIAgCg==",
					"m85znOc5znOc5znOc5znOQ==",
				},
				R: []string{
					"lQNA0DQNA0DQNA0DQNA0Dg==",
				},
			},
		},
		{
			title: "Zero divided by something",
			input: GfpolyDiv{
				A: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "Zero divided by something",
			input: GfpolyDiv{
				A: []string{
					"IAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"0AAAAAAAAAAAAAAAAAAAAA==",
					"IQAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"IAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "divisor bigger than dividend",
			input: GfpolyDiv{
				A: []string{"NeverGonnaGiveYouUpAAA=="}, B: []string{"AAAAAAAAAAAAAAAAAAAAAA==", "AAAAAAAAAAAAAAAAAAAAAA==", "NeverGonnaLetYouDownAA=="},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"NeverGonnaGiveYouUpAAA==",
				},
			},
		},
		{
			title: "rest 0",
			input: GfpolyDiv{
				A: []string{"DNWpXnnY24XecPa7a8vrEA==",
					"I8uYpCbsiPaVvUznuv1IcA==",
					"wsbiU432ARWuO93He3vbvA==",
					"zp0g3o8iNz7Y+8oUxw1vJw==",
					"J0GekE3uendpN6WUAuJ4AA==",
					"wACd0e6u1ii4AAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				}, B: []string{
					"NeverGonnaMakeYouCryAA==",
					"NeverGonnaSayGoodbyeAA==",
					"NeverGonnaTellALieAAAA==",
					"AndHurtYouAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"JAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "rest 0",
			input: GfpolyDiv{
				A: []string{
					"I20VjJmlSnRSe88gaDiLRQ==",
					"0Cw5HxJm/pfybJoQDf7/4w==",
					"8ByrMMf+vVj5r3YXUNCJ1g==",
					"rEU/f2UZRXqmZ6V7EPKfBA==",
					"LfdALhvCrdhhGZWl9l9DSg==",
					"KSUKhN0n6/DZmHPozd1prw==",
					"DQrRkuA9Zx279wAAAAAAAA==",
					"AhCEAAAAAAAAAAAAAAAAAA==",
				}, B: []string{
					"NeverGonnaMakeYouCryAA==",
					"NeverGonnaSayGoodbyeAA==",
					"NeverGonnaTellALieAAAA==",
					"AndHurtYouAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyDivExpected{
				Q: []string{
					"50AAAAAAAAAAAAAAAAAAAA==",
					"KcQAAAAAAAAAAAAAAAAAAA==",
					"DQNAAAAAAAAAAAAAAAAAAA==",
					"AhCEAAAAAAAAAAAAAAAAAA==",
				},
				R: []string{
					"AAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Q, testcase.input.Q, "Q: Expected %v, got %v", testcase.expected.Q, testcase.input.Q)
			assert.Equal(t, testcase.expected.R, testcase.input.R, "R: Expected %v, got %v", testcase.expected.R, testcase.input.R)
		})
	}
}
