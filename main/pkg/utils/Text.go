package utils

import (
	"encoding/base64"
	"math/big"
)

type Text struct {
	Content big.Int
	Len     int
}

func GetContent(s string) Text {
	byteSlice, _ := base64.StdEncoding.DecodeString(s)
	number := new(Long)
	size := (len(byteSlice) + 15) / 16
	padding := make([]byte, size*16-len(byteSlice))
	byteSlice = append(byteSlice, padding...)
	number.SetBytes(byteSlice)
	byteLen := len(byteSlice)
	if byteLen < 1 {
		byteLen = 1
	}
	return Text{Content: number.GcmToggle().Int, Len: byteLen}
}

func (t Text) Base64() string {
	return NewLongFromBigInt(t.Content).GcmToggle().GetBase64(t.Len)
}
