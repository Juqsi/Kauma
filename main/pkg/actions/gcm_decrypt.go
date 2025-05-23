package actions

import (
	"Abgabe/main/pkg/utils"
	"encoding/base64"
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

	var lastXor big.Int
	var textGcm big.Int
	var hBig big.Int

	switch args.Algorithm {
	case "aes128":
		lastXor = firstBlock(key, nonce, AesEncrypt)
		_, hBig = CalculateH(key, AesEncrypt)
		textGcm, _ = GcmBlocksEncryption(key, nonce, ciphertext, AesEncrypt)
	case "sea128":
		lastXor = firstBlock(key, nonce, Sea128Encrypt)
		_, hBig = CalculateH(key, Sea128Encrypt)
		textGcm, _ = GcmBlocksEncryption(key, nonce, ciphertext, Sea128Encrypt)
	}

	_, lBig := CalculateL(args.Ciphertext, args.Ad)

	resultGhash := GHASHBigEndian(hBig, utils.Text{Content: ciphertext, Len: (ciphertext.BitLen() + 7) / 8}, lBig, utils.Text{Content: ad, Len: (ad.BitLen() + 7) / 8})

	resultGhash = utils.NewLongFromBigInt(resultGhash).GcmToggle().Int

	tag := *resultGhash.Xor(&resultGhash, &lastXor)

	//Nur für die genaue länge
	ciphertextBytes, _ := base64.StdEncoding.DecodeString(args.Ciphertext)

	args.Plaintext = utils.NewLongFromBigInt(textGcm).GetBase64(len(ciphertextBytes))
	args.Authentic = utils.NewLongFromBigInt(tag).GetBase64(16) == args.Tag
}
