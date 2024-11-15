package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Gcm_Decrypt_Expected struct {
	Authentic bool   `json:"authentic"`
	Plaintext string `json:"plaintext"`
}

func TestGcmDecrypt_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Gcm_Decrypt
		expected Gcm_Decrypt_Expected
	}{
		{
			title: "Basic Test Encode aes",
			input: Gcm_Decrypt{
				Algorithm:  "aes128",
				Nonce:      "SwVYf3IJ2p/VTiyz",
				Key:        "bYfxz4zIS8NGWT55xSGy7Q==",
				Ciphertext: "ltCrCEYjUCHd",
				Ad:         "w8pc8/yCt2zxERPVcsnOMx8/HmrfAfGUQtD+vyMMpJ5lrF2S",
				Tag:        "d4Z2uVRSmpVE1TEa/Zhx9A==",
			},
			expected: Gcm_Decrypt_Expected{
				Authentic: true,
				Plaintext: "EREavadt90qq",
			},
		},
		{
			title: "Basic Test Encode sea",
			input: Gcm_Decrypt{
				Algorithm:  "sea128",
				Nonce:      "VOkKCCnH4EYE1z4L",
				Key:        "ByMrTiLP7isfBDL7vsKkOQ==",
				Ciphertext: "UdpDzPAafM+y",
				Ad:         "UknNF3AKBaF/8GUnFUw=",
				Tag:        "sN0+1fG+WSOHMswF7IBnZA==",
			},
			expected: Gcm_Decrypt_Expected{
				Authentic: false,
				Plaintext: "AxSiKm93Gr2+",
			},
		},
		{
			title: "failed Test Encode aes",
			input: Gcm_Decrypt{
				Algorithm:  "aes128",
				Nonce:      "50aLsYDILtAaZqbe",
				Key:        "eUisCtpuaE4sBoGAssUIHg==",
				Ciphertext: "oNA0HfQ8GdotcH47IgFBJ0NQiSLqz4GJMp+t",
				Ad:         "xhbDIZtET4WYYYsEAu6uWfVRpg8=",
				Tag:        "ovgemzM2PEV76rhFmp0SMg==",
			},
			expected: Gcm_Decrypt_Expected{
				Authentic: true,
				Plaintext: "OIRJ9HElN+HNYhaqWlk+cXwFThZDMTwXrEhd",
			},
		},
		{
			title: "failed Test Encode sea",
			input: Gcm_Decrypt{
				Algorithm:  "sea128",
				Nonce:      "50aLsYDILtAaZqbe",
				Key:        "eUisCtpuaE4sBoGAssUIHg==",
				Ciphertext: "vHh+8yTwsbK8dHTcBbVIvyQ4dl8l0jkJp4KwOMvVgw==",
				Ad:         "itzgNFd7bmvnHvc9S2PsfraTSJffPa9rBKhyTnnRGe1JUQ==",
				Tag:        "otHDGWbSZIiiJdX2D39gsg==",
			},
			expected: Gcm_Decrypt_Expected{
				Authentic: true,
				Plaintext: "5NPt2l4HX3aypuOjvRLZ+NuSX6tzwkRo15W/t7+2YQ==",
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			assert.Equal(t, testcase.expected.Authentic, testcase.input.Authentic, "Authentic: Expected %v, got %v", testcase.expected.Authentic, testcase.input.Authentic)
			assert.Equal(t, testcase.expected.Plaintext, testcase.input.Plaintext, "Plaintext: Expected %v, got %v", testcase.expected.Plaintext, testcase.input.Plaintext)
		})
	}
}
