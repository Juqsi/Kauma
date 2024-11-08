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

	var tag big.Int
	var ciphertexts big.Int
	var h big.Int
	var l big.Int

	switch args.Algorithm {

	case "aes128":
		ciphertexts, tag, h, l = GcmEncrypt(key, nonce, plaintext, ad, AesEncrypt)
	case "sea128":
		ciphertexts, tag, h, l = GcmEncrypt(key, nonce, plaintext, ad, Sea128Encrypt)
	}
	args.Ciphertext = utils.NewLongFromBigInt(ciphertexts).GetBase64(len(ciphertexts.Bytes()))
	args.Tag = utils.NewLongFromBigInt(tag).GetBase64(16)
	args.H = utils.NewLongFromBigInt(h).GetBase64(16)
	args.L = utils.NewLongFromBigInt(l).GetBase64(16)
}

func GcmEncrypt(key, nonce, plaintext, ad big.Int, encryption Encryption) (ciphertexts, tag, h, l big.Int) {
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
	ciphers := make([]big.Int, (plaintext.BitLen()+127)/128)

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
		cipherBlock = utils.NewLongFromBigInt(cipherBlock).Reverse(128).Int

		//16 Byte des Texts nehmen
		plaintextBlock := new(big.Int).And(&plaintext, sixteenByte)
		//plaintextBlock um 128 bit nach rechts shiften für die nächsten in der nächsten runde
		plaintext.Rsh(&plaintext, 128)
		// länge ermittelen in gcm
		plaintextBlockLen := (plaintextBlock.BitLen() + 7) / 8 * 8
		a := new(big.Int).SetBit(big.NewInt(0), plaintextBlockLen, 1)
		a.Sub(a, big.NewInt(1))
		cipherBlock.And(&cipherBlock, a)
		// in big endian eingetragen
		ciphers[index] = *new(big.Int).Xor(&cipherBlock, plaintextBlock)
		gcm := utils.NewLongFromBigInt(ciphers[index]).Reverse(plaintextBlockLen).Int
		// zum finalen ciphertext hinzufügen
		ciphertexts.Or(&ciphertexts, new(big.Int).Lsh(&gcm, uint(index*128)))
		index++
	}
	//Step 4
	L := big.NewInt(int64((ad.BitLen() + 7) / 8 * 8))
	L.Lsh(L, 64)
	L.Add(L, big.NewInt(int64(plaintextLen)))
	l = *L
	//l zu big Endian umdrehen
	L = &utils.NewLongFromBigInt(*L).Reverse(128).Int

	//soll 96 d0 ab
	resultGhash := GHASHBigEndian(h, ciphers, *L, ad)
	//l zu gcm umdrehen
	resultGhash = utils.NewLongFromBigInt(resultGhash).Reverse(128).Int

	tag = *resultGhash.Xor(&resultGhash, &lastXor)

	return ciphertexts, tag, h, l
}

func GHASHBigEndian(hGcm big.Int, ciphers []big.Int, l, ad big.Int) big.Int {
	sixteenByte := big.NewInt(1)
	sixteenByte.Lsh(sixteenByte, 129)
	sixteenByte.Sub(sixteenByte, big.NewInt(1))

	// tmp ist in big endian return in bigInt
	adBlock := new(big.Int).And(&ad, sixteenByte)
	//xor unnötig aber naja
	tmp := *new(big.Int).Xor(adBlock, big.NewInt(0))

	hBig := utils.NewLongFromBigInt(hGcm).Reverse(128).Int

	tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
	ad.Rsh(&ad, 128)

	for ad.BitLen() > 0 {
		adBlock = new(big.Int).And(&ad, sixteenByte)
		tmp.Xor(&tmp, adBlock)
		tmp = GfmulBigInt(tmp, *adBlock, Coeff2Number([]uint{128, 7, 2, 1, 0}))
		ad.Rsh(&ad, 128)
	}

	for _, cipherBlock := range ciphers {
		tmp.Xor(&tmp, &cipherBlock)
		tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))
	}

	tmp.Xor(&tmp, &l)
	tmp = GfmulBigInt(tmp, hBig, Coeff2Number([]uint{128, 7, 2, 1, 0}))

	return tmp
}
