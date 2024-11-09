package actions

import (
	"Abgabe/main/pkg/utils"
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
	var ciphertexts big.Int
	var h big.Int
	var l big.Int

	switch args.Algorithm {

	case "aes128":
		ciphertexts, lastXor, h, l = GcmEncrypt(key, nonce, plaintext, ad, AesEncrypt)
	case "sea128":
		ciphertexts, lastXor, h, l = GcmEncrypt(key, nonce, plaintext, ad, Sea128Encrypt)
	}

	//zu big Endian umdrehen
	lBig := &utils.NewLongFromBigInt(l).GcmToggle().Int
	ciphertextBig := utils.NewLongFromBigInt(ciphertexts).GcmToggle().Int
	hBig := utils.NewLongFromBigInt(h).GcmToggle().Int

	resultGhash := GHASHBigEndian(hBig, ciphertextBig, *lBig, ad)

	resultGhash = utils.NewLongFromBigInt(resultGhash).GcmToggle().Int

	tag := *resultGhash.Xor(&resultGhash, &lastXor)

	args.Ciphertext = utils.NewLongFromBigInt(ciphertexts).GetBase64(len(ciphertexts.Bytes()))
	args.Tag = utils.NewLongFromBigInt(tag).GetBase64(16)
	args.H = utils.NewLongFromBigInt(h).GetBase64(16)
	args.L = utils.NewLongFromBigInt(l).GetBase64(16)
}

func GcmEncrypt(key, nonce, plaintext, ad big.Int, encryption Encryption) (ciphertexts, lastXor, h, l big.Int) {
	plaintextLen := (plaintext.BitLen() + 7) / 8 * 8

	//Step 1
	nonce.Lsh(&nonce, 32).SetBit(&nonce, 0, 1)
	lastXor, err := encryption(key, nonce)
	if err != nil {
		panic("1 Error in GcmEncrypt: " + err.Error())
	}

	//Step 2
	h, err = encryption(key, *new(big.Int))
	if err != nil {
		panic("2 Error in GcmEncrypt: " + err.Error())
	}

	//Step 3
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 129)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))

	index := 0
	for plaintext.BitLen() > 0 {
		//nonce immer um 1 erhöhen
		nonce.Add(&nonce, big.NewInt(1))
		//encrypted block berechnen
		cipherBlock, err := encryption(key, nonce)
		if err != nil {
			panic("3.1 Error in GcmEncrypt: " + err.Error())
		}
		//zu big endian umdrehen für
		cipherBlock = utils.NewLongFromBigInt(cipherBlock).GcmToggle().Int

		//16 Byte des Texts nehmen
		plaintextBlock := new(big.Int).And(&plaintext, sixteenByte)

		//plaintextBlock um 128 bit nach rechts shiften für die nächsten in der nächsten runde
		plaintext.Rsh(&plaintext, 128)

		// länge ermittelen in gcm
		plaintextBlockLen := (plaintextBlock.BitLen() + 7) / 8 * 8
		a := new(big.Int).SetBit(big.NewInt(0), plaintextBlockLen, 1)
		a.Sub(a, big.NewInt(1))

		//xor
		cipherBlock = *new(big.Int).Xor(&cipherBlock, plaintextBlock)
		cipherBlock.And(&cipherBlock, a)

		// in big endian eingetragen
		gcm := utils.NewLongFromBigInt(cipherBlock).ReverseCustom(plaintextBlockLen).Int
		// zum finalen ciphertext hinzufügen
		ciphertexts.Or(&ciphertexts, new(big.Int).Lsh(&gcm, uint(index*128)))
		index++
	}

	//Step 4
	L := big.NewInt(int64((ad.BitLen() + 7) / 8 * 8))
	L.Lsh(L, 64)
	L.Add(L, big.NewInt(int64(plaintextLen)))
	l = *L

	return ciphertexts, lastXor, h, l
}

func GHASHBigEndian(hBig big.Int, ciphers big.Int, l, ad big.Int) big.Int {
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 129)
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

	tmp.Xor(&tmp, &l)
	tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
	return tmp
}
