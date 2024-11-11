package actions

import (
	"Abgabe/main/pkg/utils"
	"encoding/base64"
	"math/big"
)

type Gcm_Encrypt struct {
	Algorithm  string `json:"algorithm"`
	Nonce      string `json:"nonce"`
	Key        string `json:"key"`
	Plaintext  string `json:"plaintext"`
	Ad         string `json:"ad"`
	Ciphertext string `json:"ciphertext"`
	Tag        string `json:"tag"`
	L          string `json:"L"`
	H          string `json:"H"`
}

func (args *Gcm_Encrypt) Execute() {

	nonce := utils.NewLongFromBase64(args.Nonce).Int
	key := utils.NewLongFromBase64(args.Key).Int
	plaintext := utils.NewBigEndianLongFromGcmInBase64(args.Plaintext).Int
	ad := utils.NewBigEndianLongFromGcmInBase64(args.Ad).Int

	var lastXor big.Int
	var textGcm big.Int
	var textBig big.Int
	var hGcm big.Int
	var hBig big.Int

	switch args.Algorithm {
	case "aes128":
		lastXor = firstBlock(key, nonce, AesEncrypt)
		hGcm, hBig = calculateH(key, AesEncrypt)
		textGcm, textBig = gcmBlocksEncryption(key, nonce, plaintext, AesEncrypt)
	case "sea128":
		lastXor = firstBlock(key, nonce, Sea128Encrypt)
		hGcm, hBig = calculateH(key, Sea128Encrypt)
		textGcm, textBig = gcmBlocksEncryption(key, nonce, plaintext, Sea128Encrypt)
	}

	lGcm, lBig := calculateL(plaintext, ad)

	resultGhash := GHASHBigEndian(hBig, textBig, lBig, ad)

	resultGhash = utils.NewLongFromBigInt(resultGhash).GcmToggle().Int

	tag := *resultGhash.Xor(&resultGhash, &lastXor)

	//Nur für die genaue länge
	plaintextBytes, _ := base64.StdEncoding.DecodeString(args.Plaintext)

	args.Ciphertext = utils.NewLongFromBigInt(textGcm).GetBase64(len(plaintextBytes))
	args.Tag = utils.NewLongFromBigInt(tag).GetBase64(16)
	args.H = utils.NewLongFromBigInt(hGcm).GetBase64(16)
	args.L = utils.NewLongFromBigInt(lGcm).GetBase64(16)
}

// Step 1
func firstBlock(key, nonce big.Int, encryption Encryption) big.Int {
	nonce.Lsh(&nonce, 32).SetBit(&nonce, 0, 1)
	lastXor, err := encryption(key, nonce)
	if err != nil {
		panic("1 Error in gcmBlocksEncryption: " + err.Error())
	}
	return lastXor
}

// Step 2
func calculateH(key big.Int, encryption Encryption) (hGcm, hBig big.Int) {
	hGcm, err := encryption(key, *new(big.Int))
	if err != nil {
		panic("2 Error in gcmBlocksEncryption: " + err.Error())
	}
	hBig = utils.NewLongFromBigInt(hGcm).GcmToggle().Int
	return hGcm, hBig
}

// Step 4.3
func calculateL(plaintext, ad big.Int) (lGcm, lBig big.Int) {
	plaintextLen := (plaintext.BitLen() + 7) / 8 * 8

	lGcm = *big.NewInt(int64((ad.BitLen() + 7) / 8 * 8))
	lGcm.Lsh(&lGcm, 64)
	lGcm.Add(&lGcm, big.NewInt(int64(plaintextLen)))
	lBig = utils.NewLongFromBigInt(lGcm).GcmToggle().Int
	return lGcm, lBig
}

func gcmBlocksEncryption(key, nonce, plaintext big.Int, encryption Encryption) (textGcm, textBig big.Int) {

	//Step 3
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 128)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))

	index := 0
	for plaintext.BitLen() > 0 {
		//nonce immer um 1 erhöhen
		nonce.Add(&nonce, big.NewInt(1))
		//encrypted block berechnen
		cipherBlock, err := encryption(key, nonce)
		if err != nil {
			panic("3.1 Error in gcmBlocksEncryption: " + err.Error())
		}

		//zu big endian umdrehen für
		cipherBlock = utils.NewLongFromBigInt(cipherBlock).GcmToggle().Int

		//16 Byte des Texts nehmen
		plaintextBlock := new(big.Int).And(&plaintext, sixteenByte)

		// länge ermittelen in gcm
		plaintextBlockLen := (plaintextBlock.BitLen() + 7) / 8 * 8
		a := new(big.Int).SetBit(big.NewInt(0), plaintextBlockLen, 1)
		a.Sub(a, big.NewInt(1))

		//xor
		cipherBlock = *new(big.Int).Xor(&cipherBlock, plaintextBlock)
		cipherBlock.And(&cipherBlock, a)

		// in big endian eingetragen
		//vor veränderung für gcm zwischenspeichern
		gcm := utils.NewLongFromBigInt(cipherBlock).ReverseCustom(plaintextBlockLen).Int
		cipherBlock.Lsh(&cipherBlock, uint(index*128))
		textBig.Or(&textBig, &cipherBlock)
		//gcm eintragen beides nötig da nullen Problem
		textGcm.Lsh(&textGcm, uint(plaintextBlockLen))
		textGcm.Or(&textGcm, &gcm)

		//nächste Runde vorbereiten plaintextBlock um 128 bit nach rechts
		index++
		plaintext.Rsh(&plaintext, 128)
	}
	return textGcm, textBig
}

// gibt in bigEndian zurück
func GHASHBigEndian(hBig big.Int, ciphers big.Int, lBig, ad big.Int) big.Int {
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 128)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))

	// tmp ist in big endian return in bigInt
	adBlock := new(big.Int).And(&ad, sixteenByte)
	//xor unnötig aber naja
	tmp := *new(big.Int).Xor(adBlock, big.NewInt(0))

	tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))

	ad.Rsh(&ad, 128)

	for ad.BitLen() > 0 {
		adBlock = new(big.Int).And(&ad, sixteenByte)
		tmp.Xor(&tmp, adBlock)
		tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
		ad.Rsh(&ad, 128)
	}

	for ciphers.BitLen() > 0 {
		cipherBlock := new(big.Int).And(&ciphers, sixteenByte)
		tmp.Xor(&tmp, cipherBlock)
		tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
		ciphers.Rsh(&ciphers, 128)
	}

	tmp.Xor(&tmp, &lBig)
	tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
	return tmp
}
