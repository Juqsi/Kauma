package glasskey

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"fmt"
)

type Prng struct {
	AgencyKey string   `json:"agency_key"`
	Seed      string   `json:"seed"`
	Lengths   []int    `json:"lengths"`
	Blocks    []string `json:"blocks"`
}

func (args *Prng) Execute() {
	// Decodiere den AgencyKey und Seed von Base64
	agencyKey, err := base64.StdEncoding.DecodeString(args.AgencyKey)
	if err != nil {
		panic(fmt.Sprintf("Fehler beim Decodieren des AgencyKeys: %v", err))
	}

	seed, err := base64.StdEncoding.DecodeString(args.Seed)
	if err != nil {
		panic(fmt.Sprintf("Fehler beim Decodieren des Seeds: %v", err))
	}

	// Generiere PRNG-Blöcke
	blocks := glasskeyPRNG(agencyKey, seed, args.Lengths)

	// Blocks zuweisen
	args.Blocks = blocks
}

func glasskeyPRNG(K []byte, s []byte, lengths []int) []string {
	var result []string
	var counter uint64 = 0

	// Berechne Kstar einmal zu Beginn
	hashK := sha256.Sum256(K)
	hashS := sha256.Sum256(s)
	Kstar := append(hashK[:], hashS[:]...)

	// Erstelle ein einziges HMAC-Objekt
	h := hmac.New(sha256.New, Kstar)

	var stream []byte
	for _, length := range lengths {
		for len(stream) < length {
			iBytes := make([]byte, 8)
			binary.LittleEndian.PutUint64(iBytes, counter)

			h.Reset()
			h.Write(iBytes)
			block := h.Sum(nil)

			stream = append(stream, block...)

			counter++
		}
		encodedStream := base64.StdEncoding.EncodeToString(stream[:length])
		stream = stream[length:]
		result = append(result, encodedStream)
	}

	return result
}
