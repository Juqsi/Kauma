package actions

import (
	"Abgabe/main/pkg/utils"
	"fmt"
	"math/big"
)

type Gcm_Decrypt struct {
	Algorithm  string `json:"algorithm"`
	Nonce      string `json:"nonce"`
	Key        string `json:"key"`
	Ciphertext string `json:"ciphertext"`
	Ad         string `json:"ad"`
	Tag        string `json:"tag"`
	Authentic  bool   `json:"authentic"`
	Plaintext  string `json:"plaintext"`
}

func (args *Gcm_Decrypt) Execute() {

	nonce := utils.NewLongFromBase64(args.Nonce).Int
	key := utils.NewLongFromBase64(args.Key).Int
	ciphertext := utils.NewBigEndianLongFromGcmInBase64(args.Ciphertext).Int
	ad := utils.NewBigEndianLongFromGcmInBase64(args.Ad).Int

	var tag big.Int
	var plaintext big.Int

	switch args.Algorithm {

	case "aes128":
		plaintext, tag, _, _ = GcmEncrypt(key, nonce, ciphertext, ad, AesEncrypt)
	case "sea128":
		plaintext, tag, _, _ = GcmEncrypt(key, nonce, ciphertext, ad, Sea128Encrypt)
	}
	fmt.Println("argTag: ", args.Tag)
	fmt.Println("calcukated Tag: ", utils.NewLongFromBigInt(tag).GetBase64(16))
	args.Plaintext = utils.NewLongFromBigInt(plaintext).GetBase64(len(plaintext.Bytes()))
	args.Authentic = utils.NewLongFromBigInt(tag).GetBase64(16) == args.Tag
}
