package actions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFde_Execute(t *testing.T) {
	testcases := []struct {
		title    string
		input    Xex
		expected string
	}{
		{
			title: "Basic Test Encrypt",
			input: Xex{Mode: "encrypt", Key: "B1ygNO/CyRYIUYhTSgoUysX5Y/wWLi4UiWaVeloUWs0=",
				Tweak: "6VXORr+YYHrd2nVe0OlA+Q==", Input: "/aOg4jMocLkBLkDLgkHYtFKc2L9jjyd2WXSSyxXQikpMY9ZRnsJE76e9dW9olZIW"},
			expected: "mHAVhRCKPAPx0BcufG5BZ4+/CbneMV/gRvqK5rtLe0OJgpDU5iT7z2P0R7gEeRDO",
		},
		{
			title: "Basic Test Decrypt",
			input: Xex{Mode: "decrypt", Key: "B1ygNO/CyRYIUYhTSgoUysX5Y/wWLi4UiWaVeloUWs0=",
				Tweak: "6VXORr+YYHrd2nVe0OlA+Q==", Input: "lr/ItaYGFXCtHhdPndE65yg7u/GIdM9wscABiiFOUH2Sbyc2UFMlIRSMnZrYCW1a"},
			expected: "SGV5IHdpZSBrcmFzcyBkYXMgZnVua3Rpb25pZXJ0IGphIG9mZmVuYmFyIGVjaHQu",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.title, func(t *testing.T) {
			testcase.input.Execute()
			response := testcase.input.Result
			assert.Equal(t, testcase.expected, response, "Expected %v, got %v", testcase.expected, response)
		})
	}
}
