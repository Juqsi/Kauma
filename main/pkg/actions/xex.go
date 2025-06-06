package actions

import (
	"Abgabe/main/pkg/utils"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"math/big"
)

type Xex struct {
	Mode   string `json:"mode"`
	Key    string `json:"key"`
	Tweak  string `json:"tweak"`
	Input  string `json:"input"`
	Result string `json:"result"`
}

func (f *Xex) Execute() {
	fullKey := utils.NewLongFromBase64(f.Key).Int.Bytes()
	if len(fullKey) != 32 {
		f.Result = "Invalid key"
		return
	}
	key1 := *new(big.Int).SetBytes(fullKey[:16])
	key2 := *new(big.Int).SetBytes(fullKey[16:])

	tweak := utils.NewLongFromBase64(f.Tweak).Int
	//Step 1
	encryptedTweak, err := Sea128Encrypt(key2, tweak)
	if err != nil {
		f.Result = "tweak encryption error"
		return
	}

	var text []byte
	input, err := base64.StdEncoding.DecodeString(f.Input)
	if err != nil {
		panic(fmt.Sprintf("Error invalid input: %+v, %s\n", *f, err.Error()))
	}
	if f.Mode == "encrypt" {
		text, err = FdeXexEncrypt(key1, encryptedTweak, input)
	} else if f.Mode == "decrypt" {
		text, err = FdeXexDecrypt(key1, encryptedTweak, input)
	} else {
		panic(fmt.Sprintf("Error invalid mode: %+v, %s\n", *f, err.Error()))
	}
	f.Result = base64.StdEncoding.EncodeToString(text)
}

func FdeXexEncrypt(key, tweak big.Int, message []byte) (cipher []byte, err error) {
	blocks := getBlocks(message, 16)
	a := Coeff2Number([]uint{1})
	for _, block := range blocks {
		//Step 2
		block.Xor(&block, &tweak)
		//Step 3
		encryptedBlock, err := Sea128Encrypt(key, block)
		if err != nil {
			return cipher, err
		}
		//Step4
		encryptedBlock.Xor(&encryptedBlock, &tweak)
		cipher = append(cipher, utils.NewLongFromBigInt(encryptedBlock).Bytes(aes.BlockSize)...)
		//Step 5
		tweak = Gfmul128(*new(big.Int).SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian()), a)
		tweak.SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian())
	}
	return cipher, err
}

func FdeXexDecrypt(key, tweak big.Int, cipher []byte) (text []byte, err error) {
	blocks := getBlocks(cipher, 16)
	a := Coeff2Number([]uint{1})
	for _, block := range blocks {
		//Step 2
		block.Xor(&block, &tweak)

		//Step 3
		decryptedBlock, err := Sea128Decrypt(key, block)
		if err != nil {
			return text, err
		}
		//Step 4
		decryptedBlock.Xor(&decryptedBlock, &tweak)
		text = append(text, utils.NewLongFromBigInt(decryptedBlock).Bytes(aes.BlockSize)...)

		//Step 5
		tweak = Gfmul128(*new(big.Int).SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian()), a)
		tweak.SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian())
	}
	return text, err
}

func getBlocks(message []byte, size int) []big.Int {
	blocks := *new([]big.Int)
	for len(message) > 0 {
		blocks = append(blocks, *new(big.Int).SetBytes(message[:size]))
		message = message[size:]
	}
	return blocks
}
