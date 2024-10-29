package actions

import (
	"Abgabe/main/pkg/utils"
	"crypto/aes"
	"math/big"
)

type sea128 struct {
	Mode   string `json:"Mode"`
	Key    string `json:"Key"`
	Input  string `json:"Input"`
	Result string
}

func (s *sea128) Execute() {
	var a *big.Int
	var err error

	seaConstant, _ := new(big.Int).SetString("c0ffeec0ffeec0ffeec0ffeec0ffee11", 16)

	message := &utils.NewLongFromBase64(s.Input).Int
	key := &utils.NewLongFromBase64(s.Key).Int

	if s.Mode == "encrypt" {
		a, err = Sea128Encrypt(key, message, seaConstant)
	} else if s.Mode == "decrypt" {
		a, err = Sea128Decrypt(key, message, seaConstant)
	} else {
		s.Result = "Invalid mode"
		return
	}

	if err != nil {
		s.Result = err.Error()
		return
	}
	s.Result = utils.NewLongFromBigInt(a).GetBase64(16)
}

func Sea128Encrypt(key, message, seaConstant *big.Int) (*big.Int, error) {
	cipher, err := AesEncrypt(key, message)
	if err != nil {
		return &big.Int{}, err
	}
	cipher.Xor(cipher, seaConstant)

	return cipher, nil
}

func AesEncrypt(key, message *big.Int) (*big.Int, error) {
	block, err := aes.NewCipher(key.Bytes())
	if err != nil {
		return &big.Int{}, err
	}

	ciphertext := make([]byte, aes.BlockSize)
	block.Encrypt(ciphertext, message.Bytes())

	return new(big.Int).SetBytes(ciphertext), nil
}

func Sea128Decrypt(key, ciphertext, seaConstant *big.Int) (*big.Int, error) {
	ciphertext.Xor(ciphertext, seaConstant)
	plaintext, err := AesDecrypt(key, ciphertext)
	if err != nil {
		return &big.Int{}, err
	}

	return plaintext, nil
}

func AesDecrypt(key, ciphertext *big.Int) (*big.Int, error) {
	block, err := aes.NewCipher(key.Bytes())
	if err != nil {
		return &big.Int{}, err
	}

	plaintext := make([]byte, aes.BlockSize)
	block.Decrypt(plaintext, ciphertext.Bytes())

	return new(big.Int).SetBytes(plaintext), nil
}
