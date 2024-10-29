package utils

import (
	"encoding/base64"
	"math/big"
)

type Long struct {
	big.Int
}

func NewLongFromBigInt(b *big.Int) *Long {
	long := new(Long)
	long.Set(b)
	return long
}

func NewLongFromLittleEndianInBase64(s string) *Long {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := new(Long)
	number.SetBytes(byteSlice)
	number.SetBytes(number.GetLittleEndian())
	return number
}

func NewLongFromBase64(s string) *Long {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := new(Long)
	number.SetBytes(byteSlice)
	return number
}

func (number *Long) GetLittleEndian() []byte {
	byteSlice := make([]byte, len(number.Bytes()))
	for i, value := range number.Bytes() {
		byteSlice[len(byteSlice)-1-i] = value
	}
	return byteSlice
}

func (number *Long) GetLittleEndianInBase64(minLenght int) string {
	numberBytes := number.GetLittleEndian()
	for i := len(numberBytes); i < minLenght; i++ {
		numberBytes = append(numberBytes, 0)
	}
	return base64.StdEncoding.EncodeToString(numberBytes)
}

func (number *Long) GetBase64(minLenght int) string {
	numberBytes := number.Int.Bytes()
	for i := len(numberBytes); i < minLenght; i++ {
		numberBytes = append(numberBytes, 0)
	}
	return base64.StdEncoding.EncodeToString(numberBytes)
}
