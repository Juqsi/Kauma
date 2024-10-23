package utils

import (
	"encoding/base64"
	"math/big"
)

type Long struct {
	big.Int
}

func NewLong(size int) *Long {
	buff := make([]byte, size)
	a := new(Long)
	a.SetBytes(buff)
	return a
}

func NewLongFromBigInt(b *big.Int) *Long {
	long := new(Long)
	long.Set(b)
	return long
}

func NewLongFromBase64InBigEndian(s string) *Long {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := NewLong(len(byteSlice))
	number.SetBytes(byteSlice)
	number.SetBytes(number.GetBigEndian())
	return number
}

func (number *Long) GetBigEndian() []byte {
	byteSlice := make([]byte, len(number.Bytes()))
	for i, value := range number.Bytes() {
		byteSlice[len(byteSlice)-1-i] = value
	}
	return byteSlice
}

func (number *Long) GetBigEndianInBase64() string {
	return base64.StdEncoding.EncodeToString(number.GetBigEndian())
}

func (number *Long) BigInt() *big.Int {
	return new(big.Int).SetBytes(number.Bytes())
}
