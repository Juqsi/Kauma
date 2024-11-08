package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Gcm_Encrypt_Expected struct {
	Ciphertext string `json:"ciphertext"`
	Tag        string `json:"tag"`
	L          string `json:"L"`
	H          string `json:"H"`
}

func TestGcmEncrypt_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Gcm_Encrypt
		expected Gcm_Encrypt_Expected
	}{
		{
			title: "Basic Test Encode aes",
			input: Gcm_Encrypt{
				Algorithm: "aes128",
				Nonce:     "4gF+BtR3ku/PUQci",
				Key:       "Xjq/GkpTSWoe3ZH0F+tjrQ==",
				Plaintext: "RGFzIGlzdCBlaW4gVGVzdA==",
				Ad:        "QUQtRGF0ZW4="},
			expected: Gcm_Encrypt_Expected{
				Ciphertext: "ET3RmvH/Hbuxba63EuPRrw==",
				Tag:        "Mp0APJb/ZIURRwQlMgNN/w==",
				L:          "AAAAAAAAAEAAAAAAAAAAgA==",
				H:          "Bu6ywbsUKlpmZXMQyuGAng=="},
		},
		{
			title: "Basic Test Encode sea",
			input: Gcm_Encrypt{
				Algorithm: "sea128",
				Nonce:     "4gF+BtR3ku/PUQci",
				Key:       "Xjq/GkpTSWoe3ZH0F+tjrQ==",
				Plaintext: "RGFzIGlzdCBlaW4gVGVzdA==",
				Ad:        "QUQtRGF0ZW4="},
			expected: Gcm_Encrypt_Expected{
				Ciphertext: "0cI/Wg4R3URfrVFZ0hw/vg==",
				Tag:        "ysDdzOSnqLH0MQ+Mkb23gw==",
				L:          "AAAAAAAAAEAAAAAAAAAAgA==",
				H:          "xhFcAUT66qWIpYz+Ch5ujw=="},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.H, testcase.input.H, "H: Expected %v, got %v", testcase.expected.H, testcase.input.H)
			assert.Equal(t, testcase.expected.Ciphertext, testcase.input.Ciphertext, "Ciphertext: Expected %v, got %v", testcase.expected.Ciphertext, testcase.input.Ciphertext)
			assert.Equal(t, testcase.expected.L, testcase.input.L, "L: Expected %v, got %v", testcase.expected.L, testcase.input.L)
			assert.Equal(t, testcase.expected.Tag, testcase.input.Tag, "Tag: Expected %v, got %v", testcase.expected.Tag, testcase.input.Tag)
		})
	}
}
