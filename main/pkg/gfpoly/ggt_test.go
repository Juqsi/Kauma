package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyGgtExpected struct {
	G []string `json:"G"`
}

func TestGfpolyGgf_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyGgt
		expected GfpolyGgtExpected
	}{
		{
			title: "Basic Test Task",
			input: GfpolyGgt{
				A: []string{
					"DNWpXnnY24XecPa7a8vrEA==",
					"I8uYpCbsiPaVvUznuv1IcA==",
					"wsbiU432ARWuO93He3vbvA==",
					"zp0g3o8iNz7Y+8oUxw1vJw==",
					"J0GekE3uendpN6WUAuJ4AA==",
					"wACd0e6u1ii4AAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"I20VjJmlSnRSe88gaDiLRQ==",
					"0Cw5HxJm/pfybJoQDf7/4w==",
					"8ByrMMf+vVj5r3YXUNCJ1g==",
					"rEU/f2UZRXqmZ6V7EPKfBA==",
					"LfdALhvCrdhhGZWl9l9DSg==",
					"KSUKhN0n6/DZmHPozd1prw==",
					"DQrRkuA9Zx279wAAAAAAAA==",
					"AhCEAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyGgtExpected{
				G: []string{
					"NeverGonnaMakeYouCryAA==",
					"NeverGonnaSayGoodbyeAA==",
					"NeverGonnaTellALieAAAA==",
					"AndHurtYouAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
		{
			title: "Basic Test Equal",
			input: GfpolyGgt{
				A: []string{
					"DNWpXnnY24XecPa7a8vrEA==",
					"I8uYpCbsiPaVvUznuv1IcA==",
					"wsbiU432ARWuO93He3vbvA==",
					"zp0g3o8iNz7Y+8oUxw1vJw==",
					"J0GekE3uendpN6WUAuJ4AA==",
					"wACd0e6u1ii4AAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
				B: []string{
					"DNWpXnnY24XecPa7a8vrEA==",
					"I8uYpCbsiPaVvUznuv1IcA==",
					"wsbiU432ARWuO93He3vbvA==",
					"zp0g3o8iNz7Y+8oUxw1vJw==",
					"J0GekE3uendpN6WUAuJ4AA==",
					"wACd0e6u1ii4AAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
			},
			expected: GfpolyGgtExpected{
				G: []string{
					"DNWpXnnY24XecPa7a8vrEA==",
					"I8uYpCbsiPaVvUznuv1IcA==",
					"wsbiU432ARWuO93He3vbvA==",
					"zp0g3o8iNz7Y+8oUxw1vJw==",
					"J0GekE3uendpN6WUAuJ4AA==",
					"wACd0e6u1ii4AAAAAAAAAA==",
					"ACAAAAAAAAAAAAAAAAAAAA==",
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.G, testcase.input.G, "Expected \n %v\n, got\n %v", testcase.expected.G, testcase.input.G)
		})
	}
}
