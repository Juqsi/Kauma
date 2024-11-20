package actions

import (
	"Abgabe/main/pkg/utils"
	"crypto/aes"
	"fmt"
	"math/big"
)

type Sea128 struct {
	Mode   string `json:"Mode"`
	Key    string `json:"Key"`
	Input  string `json:"Input"`
	Result string
}

type Encryption func(key, message big.Int) (big.Int, error)
type Decryption func(key, message big.Int) (big.Int, error)

func (s *Sea128) Execute() {
	var a big.Int
	var err error

	message := utils.NewLongFromBase64(s.Input).Int
	key := utils.NewLongFromBase64(s.Key).Int

	if s.Mode == "encrypt" {
		a, err = Sea128Encrypt(key, message)
	} else if s.Mode == "decrypt" {
		a, err = Sea128Decrypt(key, message)
	} else {
		panic(fmt.Sprintf("Error invalid mode: %+v\n", *s))
	}

	if err != nil {
		panic(fmt.Sprintf("Error: %+v, %s\n", *s, err.Error()))
		return
	}

	s.Result = utils.NewLongFromBigInt(a).GetBase64(16)
}

func Sea128Encrypt(key, message big.Int) (big.Int, error) {
	seaConstant, _ := new(big.Int).SetString("c0ffeec0ffeec0ffeec0ffeec0ffee11", 16)

	cipher, err := AesEncrypt(key, message)
	if err != nil {
		return big.Int{}, err
	}
	cipher.Xor(&cipher, seaConstant)

	return cipher, nil
}

func AesEncrypt(key, message big.Int) (big.Int, error) {
	block, err := aes.NewCipher(utils.NewLongFromBigInt(key).Bytes(aes.BlockSize))
	if err != nil {
		return big.Int{}, err
	}

	mes := utils.NewLongFromBigInt(message).Bytes(aes.BlockSize)
	ciphertext := make([]byte, aes.BlockSize)
	block.Encrypt(ciphertext, mes)

	return *new(big.Int).SetBytes(ciphertext), nil
}

func Sea128Decrypt(key, ciphertext big.Int) (big.Int, error) {
	seaConstant, _ := new(big.Int).SetString("c0ffeec0ffeec0ffeec0ffeec0ffee11", 16)

	ciphertext1 := new(big.Int).Xor(&ciphertext, seaConstant)
	plaintext, err := AesDecrypt(key, *ciphertext1)
	if err != nil {
		return big.Int{}, err
	}

	return plaintext, nil
}

func AesDecrypt(key, ciphertext big.Int) (big.Int, error) {
	block, err := aes.NewCipher(utils.NewLongFromBigInt(key).Bytes(aes.BlockSize))
	if err != nil {
		return big.Int{}, err
	}

	cip := utils.NewLongFromBigInt(ciphertext).Bytes(aes.BlockSize)
	plaintext := make([]byte, aes.BlockSize)
	block.Decrypt(plaintext, cip)

	return *new(big.Int).SetBytes(plaintext), nil
}
