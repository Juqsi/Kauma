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

// Returns a Long from a base64 encoded string in Bigendian
func NewBigEndianLongFromGcmInBase64(s string) *Long {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := new(Long)
	size := (len(byteSlice) + 15) / 16
	padding := make([]byte, size*16-len(byteSlice))
	byteSlice = append(byteSlice, padding...)
	number.SetBytes(byteSlice)
	return number.Reverse(len(byteSlice) * 8)
}

func NewLongFromBase64(s string) *Long {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := new(Long)
	number.SetBytes(byteSlice)
	return number
}

func (number *Long) GetLittleEndian() []byte {
	byteSlice := make([]byte, 16)
	for i, value := range number.Bytes(16) {
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
	numberBytes := number.Bytes(length)
	for i := len(numberBytes); i < length; i++ {
		numberBytes = append(numberBytes, 0)
	}
	return base64.StdEncoding.EncodeToString(numberBytes[:length])
}

func (number *Long) Bytes(length int) []byte {
	return append(make([]byte, length-len(number.Int.Bytes())), number.Int.Bytes()...)[:length]

}

func Xor(a, b []byte) []byte {
	lenC := len(a)
	if len(b) > lenC {
		lenC = len(b)
	}
	c := make([]byte, lenC)
	for j := 0; j < lenC; j++ {
		c[j] = a[j] ^ b[j]
	}
	return c
}

// Wenn gcm to bigendian dann muss number in gcm mit 0 gepaddet sein
// Wenn big endian zu gcm dann wird number in gcm mit 0 gepaddet sein, dadurch einfach links mit 0 auffüllen
// Bei gcm hat weiß man dann wie viele 0 links fehlen also bis zum lsb
// Bei big endian weiß man wir vielen 0 links fehlen also bis zum msb
func (number *Long) Reverse(bitLen int) *Long {
	reversed := big.NewInt(0)
	for i := 0; i < bitLen; i++ {
		if number.Bit(i) == 1 {
			reversed.SetBit(reversed, bitLen-1-i, 1)
		}
	}

	return NewLongFromBigInt(*reversed)
}
