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
		hGcm, hBig = CalculateH(key, AesEncrypt)
		textGcm, textBig = GcmBlocksEncryption(key, nonce, plaintext, AesEncrypt)
	case "sea128":
		lastXor = firstBlock(key, nonce, Sea128Encrypt)
		hGcm, hBig = CalculateH(key, Sea128Encrypt)
		textGcm, textBig = GcmBlocksEncryption(key, nonce, plaintext, Sea128Encrypt)
	}

	lGcm, lBig := CalculateL(args.Plaintext, args.Ad)

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
		panic("1 Error in GcmBlocksEncryption: " + err.Error())
	}
	return lastXor
}

// Step 2
func CalculateH(key big.Int, encryption Encryption) (hGcm, hBig big.Int) {
	hGcm, err := encryption(key, *new(big.Int))
	if err != nil {
		panic("2 Error in GcmBlocksEncryption: " + err.Error())
	}
	hBig = utils.NewLongFromBigInt(hGcm).GcmToggle().Int
	return hGcm, hBig
}

// Step 4.3
func CalculateL(encodedPlaintext, encodedAd string) (lGcm, lBig big.Int) {
	plaintext, err := base64.StdEncoding.DecodeString(encodedPlaintext)
	if err != nil {
		panic("base64 decode error")

	}

	ad, err := base64.StdEncoding.DecodeString(encodedAd)
	if err != nil {
		panic("base64 decode error")
	}

	plaintextLen := len(plaintext) * 8
	adLen := len(ad) * 8

	lGcm = *big.NewInt(int64(adLen))
	lGcm.Lsh(&lGcm, 64)
	lGcm.Add(&lGcm, big.NewInt(int64(plaintextLen)))

	lBig = utils.NewLongFromBigInt(lGcm).GcmToggle().Int
	return lGcm, lBig
}

func GcmBlocksEncryption(key, nonce, plaintxt big.Int, encryption Encryption) (textGcm, textBig big.Int) {
	plaintext := *new(big.Int).Set(&plaintxt)

	//Step 3
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 128)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))

	index := 0
	for plaintext.Sign() > 0 {
		//nonce immer um 1 erhöhen
		nonce.Add(&nonce, big.NewInt(1))
		//encrypted block berechnen
		cipherBlock, err := encryption(key, nonce)
		if err != nil {
			panic("3.1 Error in GcmBlocksEncryption: " + err.Error())
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
func GHASHBigEndian(hBig big.Int, ciphersText big.Int, lBig, AssociatedData big.Int) big.Int {
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 128)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))
	ad := *new(big.Int).Set(&AssociatedData)
	ciphers := *new(big.Int).Set(&ciphersText)

	// tmp ist in big endian return in bigInt
	adBlock := new(big.Int).And(&ad, sixteenByte)
	//xor unnötig aber naja
	tmp := *new(big.Int).Xor(adBlock, big.NewInt(0))

	tmp = Gfmul128(tmp, hBig)

	ad.Rsh(&ad, 128)

	for ad.Sign() > 0 {
		adBlock = new(big.Int).And(&ad, sixteenByte)
		tmp.Xor(&tmp, adBlock)
		tmp = Gfmul128(tmp, hBig)
		ad.Rsh(&ad, 128)
	}

	for ciphers.BitLen() > 0 {
		cipherBlock := new(big.Int).And(&ciphers, sixteenByte)
		tmp.Xor(&tmp, cipherBlock)
		tmp = Gfmul128(tmp, hBig)
		ciphers.Rsh(&ciphers, 128)
	}

	tmp.Xor(&tmp, &lBig)
	tmp = Gfmul128(tmp, hBig)
	return tmp
}
