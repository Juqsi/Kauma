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
