package utils

import (
	"encoding/base64"
	"math/big"
)

type Long struct {
	big.Int
}

func NewLongFromBigInt(b big.Int) *Long {
	long := new(Long)
	long.Set(&b)
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
	byteSlice := make([]byte, len(number.Int.Bytes()))
	for i, value := range number.Int.Bytes() {
		byteSlice[len(byteSlice)-1-i] = value
	}
	return byteSlice
}

func (number *Long) GetLittleEndianInBase64(length int) string {
	numberBytes := number.GetLittleEndian()
	for i := len(numberBytes); i < length; i++ {
		numberBytes = append(numberBytes, 0)
	}
	return base64.StdEncoding.EncodeToString(numberBytes[:length])
}

func (number *Long) GetBase64(length int) string {
	numberBytes := number.Int.Bytes()
	for i := len(numberBytes); i < length; i++ {
		numberBytes = append(numberBytes, 0)
	}
	return base64.StdEncoding.EncodeToString(numberBytes[:length])
}

func (number *Long) Bytes(length int) []byte {
	return append(make([]byte, length-len(number.Int.Bytes())), number.Int.Bytes()...)[:length]

}
