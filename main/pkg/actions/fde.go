package actions

import (
	"Abgabe/main/pkg/utils"
	"encoding/base64"
	"math/big"
)

type fde struct {
	Mode   string `json:"mode"`
	Key    string `json:"key"`
	Tweak  string `json:"tweak"`
	Input  string `json:"input"`
	Result string `json:"result"`
}

func (f *fde) Execute() {
	fullKey := utils.NewLongFromBase64(f.Key).Int.Bytes()
	if len(fullKey) != 32 {
		f.Result = "Invalid key"
		return
	}
	key1 := new(big.Int).SetBytes(fullKey[:16])
	key2 := new(big.Int).SetBytes(fullKey[16:])

	seaConst, _ := new(big.Int).SetString("c0ffeec0ffeec0ffeec0ffeec0ffee11", 16)

	tweak := &utils.NewLongFromBase64(f.Tweak).Int
	//Step 1
	encryptedTweak, err := Sea128Encrypt(key2, tweak, seaConst)
	if err != nil {
		f.Result = "tweak encryption error"
		return
	}

	var text []byte
	input, err := base64.StdEncoding.DecodeString(f.Input)
	if err != nil {
		f.Result = "Invalid input"
		return
	}
	if f.Mode == "encrypt" {
		text, err = FdeXexEncrypt(key1, encryptedTweak, seaConst, input)
	} else if f.Mode == "decrypt" {
		text, err = FdeXexDecrypt(key1, encryptedTweak, seaConst, input)
	} else {
		f.Result = "Invalid mode"
		return
	}
	f.Result = base64.StdEncoding.EncodeToString(text)
}

func FdeXexEncrypt(key, tweak, seaConst *big.Int, message []byte) (cipher []byte, err error) {
	blocks := getBlocks(message, 16)
	a := Coeff2Number([]uint{1})
	for _, block := range blocks {
		//Step 2
		block.Xor(block, tweak)
		//Step 3
		encryptedBlock, err := Sea128Encrypt(key, block, seaConst)
		if err != nil {
			return cipher, err
		}
		//Step4
		encryptedBlock.Xor(encryptedBlock, tweak)
		cipher = append(cipher, encryptedBlock.Bytes()...)
		//Step 5
		tweak = GfmulBigInt(new(big.Int).SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian()), a, Coeff2Number([]uint{128, 7, 2, 1, 0}))
		tweak.SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian())
	}
	return cipher, err
}

func FdeXexDecrypt(key, tweak, seaConst *big.Int, cipher []byte) (text []byte, err error) {
	blocks := getBlocks(cipher, 16)
	a := Coeff2Number([]uint{1})
	for _, block := range blocks {
		//Step 2
		block.Xor(block, tweak)

		//Step 3
		decryptedBlock, err := Sea128Decrypt(key, block, seaConst)
		if err != nil {
			return text, err
		}
		//Step 4
		decryptedBlock.Xor(decryptedBlock, tweak)
		text = append(text, decryptedBlock.Bytes()...)

		//Step 5
		tweak = GfmulBigInt(new(big.Int).SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian()), a, Coeff2Number([]uint{128, 7, 2, 1, 0}))
		tweak.SetBytes(utils.NewLongFromBigInt(tweak).GetLittleEndian())
	}
	return text, err
}

func getBlocks(message []byte, size int) []*big.Int {
	blocks := *new([]*big.Int)
	for len(message) > 0 {
		blocks = append(blocks, new(big.Int).SetBytes(message[:size]))
		message = message[size:]
	}
	return blocks
}
