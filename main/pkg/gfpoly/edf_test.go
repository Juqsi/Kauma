package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GfpolyEdfExpected struct {
	Factors [][]string
}

func TestGfpolyEdf_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GfpolyEdf
		expected GfpolyEdfExpected
	}{
		{
			title: "Basic Test Task",
			input: GfpolyEdf{
				F: []string{
					"mmAAAAAAAAAAAAAAAAAAAA==",
					"AbAAAAAAAAAAAAAAAAAAAA==",
					"zgAAAAAAAAAAAAAAAAAAAA==",
					"FwAAAAAAAAAAAAAAAAAAAA==",
					"AAAAAAAAAAAAAAAAAAAAAA==",
					"wAAAAAAAAAAAAAAAAAAAAA==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
				D: 3,
			},
			expected: GfpolyEdfExpected{
				Factors: [][]string{
					{
						"iwAAAAAAAAAAAAAAAAAAAA==",
						"CAAAAAAAAAAAAAAAAAAAAA==",
						"AAAAAAAAAAAAAAAAAAAAAA==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
					},
					{
						"kAAAAAAAAAAAAAAAAAAAAA==",
						"CAAAAAAAAAAAAAAAAAAAAA==",
						"wAAAAAAAAAAAAAAAAAAAAA==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
					},
				},
			},
		},
		{
			title: "testcase 024",
			input: GfpolyEdf{
				F: []string{
					"mHM50z8mq7Btg+2wRK+nlQ==",
					"Ghd/G7qfdoIbC62BFdGVlw==",
					"rl1KedH8n8edUh2waP+gNw==",
					"RR5n7GuNMyEYQE09EwLhNw==",
					"fFiek8RDjGtD7mhx4w5tEw==",
					"1cpb/J8ddQ5ZzIe+kb2FTQ==",
					"2xEVzglpTOzUurPTnPiWnA==",
					"tDMAOzidtAhWUVnN0Sr70w==",
					"KwSWLNRPEQNPFKefvn9YoA==",
					"A0KjF+OSKb9t2DgCVzwKTw==",
					"cf9LZXtG8GsM3TKzSdI0Lw==",
					"12EGf5vuJO6EfQhxrw+dCQ==",
					"gAAAAAAAAAAAAAAAAAAAAA==",
				},
				D: 6,
			},
			expected: GfpolyEdfExpected{
				Factors: [][]string{
					{
						"0x/jlB1iwCmq/WzR8GzdPg==",
						"umkVa32vrvUQBI4Rkj7z8A==",
						"N9VefMIthuWwbgUvYCGRxA==",
						"HDyQpfHyQUsR9bpmu3VUPw==",
						"8A/5HorZhK2boCz8jjCJYA==",
						"6Oe207CqJoJkAStsl+zczA==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
					},
					{
						"bjdQdlkoJIchw2qbblYFpg==",
						"wyWWdTEa94/IQLMnW9INhA==",
						"6jFELK+bftALDGPn10nGaw==",
						"hHbtvRAiVXyuJCZVkQSsCw==",
						"eSeLPN5XBkKLQT3bv7iiBg==",
						"P4awrCtEAmzgfCMdOONBxQ==",
						"gAAAAAAAAAAAAAAAAAAAAA==",
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
