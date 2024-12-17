package gfpoly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GcmCrackExpected struct {
	Tag  string
	H    string
	Mask string
}

func TestGcmCrack_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    GcmCrack
		expected GcmCrackExpected
	}{
		{
			title: "Basic Test Task",
			input: GcmCrack{
				Nonce: "4gF+BtR3ku/PUQci",
				M1: message{
					Ciphertext:     "CGOkZDnJEt24aVV8mqQq+P4pouVDWhAYj0SN5MDAgg==",
					AssociatedData: "TmFjaHJpY2h0IDE=",
					Tag:            "GC9neV3aZLnmznTIWqCC4A==",
				},
				M2: message{
					Ciphertext:     "FnWyLSTfRrO8Y1MuhLIs6A==",
					AssociatedData: "",
					Tag:            "gb2ph1vzwU85/FsUg51t3Q==",
				},
				M3: message{
					Ciphertext:     "CGOkZDnJEt25aV58iaMt6O8+8chKVh0Eg1XFxA==",
					AssociatedData: "TmFjaHJpY2h0IDM=",
					Tag:            "+/aDjsAzTseDLuM4jt5Q6Q==",
				},
				Forgery: Forgery{
					Ciphertext:     "AXe/ZQ==",
					AssociatedData: "",
				},
			},
			expected: GcmCrackExpected{
				Tag:  "Y16EEEO1sgJX3IsJSwEXlA==",
				H:    "Nxn7h7ruk8eiNAG6AfhUFg==",
				Mask: "tXjFK5vCqIPl6fKAJAyy9A==",
			},
		}, {
			title: "TC004",
			input: GcmCrack{
				Nonce: "YboP6dIQFdu31NLj",
				M1: message{
					Ciphertext:     "AA==",
					AssociatedData: "",
					Tag:            "j08k1qlTnc4DG9GNYe5LMA==",
				},
				M2: message{
					Ciphertext:     "4YWg3Ak7ehMjgL/lHC+h0LJx",
					AssociatedData: "sfJmDicCC5FFAqS2k/Il",
					Tag:            "Y841Gau407BeWRZwWohgrw==",
				},
				M3: message{
					Ciphertext:     "X8buN7x+6fy4Ow==",
					AssociatedData: "",
					Tag:            "HItl3dOadXt3E3KWGzSnyA==",
				},
				Forgery: Forgery{
					Ciphertext:     "UXXMhLXP8XF7",
					AssociatedData: "",
				},
			},
			expected: GcmCrackExpected{
				Tag:  "6/z9g/O0KhSFtFBskx3zhA==",
				H:    "L6qefn4OL06sCTP0vx3Gvg==",
				Mask: "pm0a9I45a4LGZJjPA0i3DQ==",
			},
		},
		{
			title: "TC007",
			input: GcmCrack{
				Nonce: "z9Rd0XIJkW0R9wkP",
				M1: message{
					Ciphertext:     "PVFEuwdtKsIViPKt4vMkO0W2X+08cmPRolRufIYamw==",
					AssociatedData: "ovwp5DM=",
					Tag:            "wbZ4DoZGShdEY7oUtG7txw==",
				},
				M2: message{
					Ciphertext:     "9cQKidTLnmA9x3iQg9I=",
					AssociatedData: "N83i/C+WNnI6y41w+igAPQA=",
					Tag:            "XnTTDY6uafdm4TCLgXywUw==",
				},
				M3: message{
					Ciphertext:     "EFipXcz6U8oNoh1FwgASo0PH/YpE0mcTnA==",
					AssociatedData: "",
					Tag:            "qnI1ZSfESt435kHMllTU1Q==",
				},
				Forgery: Forgery{
					Ciphertext:     "YzdYTe7WKj8ZqhRg0pC2dwPlEQ==",
					AssociatedData: "",
				},
			},
			expected: GcmCrackExpected{
				Tag:  "E1F9Bo5h5Nk/o8KG3IG0MQ==",
				H:    "K2V51rMT9BvPWLrvttOqNQ==",
				Mask: "DwEnwF8sT2XVN6iKwiw5DQ==",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Tag, testcase.input.Tag, "Tag: Expected \n %v\n, got\n %v", testcase.expected.Tag, testcase.input.Tag)
			assert.Equal(t, testcase.expected.Mask, testcase.input.Mask, "Mask: Expected \n %v\n, got\n %v", testcase.expected.Mask, testcase.input.Mask)
			assert.Equal(t, testcase.expected.H, testcase.input.H, "H: Expected \n %v\n, got\n %v", testcase.expected.H, testcase.input.H)
		})
	}
}
