package actions

import (
	"Abgabe/main/pkg/utils"
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
	var plaintext big.Int
	var h big.Int
	var l big.Int

	switch args.Algorithm {

	case "aes128":
		plaintext, lastXor, h, l = GcmEncrypt(key, nonce, ciphertext, ad, AesEncrypt)
	case "sea128":
		plaintext, lastXor, h, l = GcmEncrypt(key, nonce, ciphertext, ad, Sea128Encrypt)
	}

	//zu big Endian umdrehen
	lBig := &utils.NewLongFromBigInt(l).GcmToggle().Int
	hBig := utils.NewLongFromBigInt(h).GcmToggle().Int

	resultGhash := GHASHBigEndian(hBig, ciphertext, *lBig, ad)

	resultGhash = utils.NewLongFromBigInt(resultGhash).GcmToggle().Int

	tag := *resultGhash.Xor(&resultGhash, &lastXor)

	args.Plaintext = utils.NewLongFromBigInt(plaintext).GetBase64(len(plaintext.Bytes()))
	args.Authentic = utils.NewLongFromBigInt(tag).GetBase64(16) == args.Tag
}
