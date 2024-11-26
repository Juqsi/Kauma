package actions

import (
	"Abgabe/main/pkg/utils"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

const PADDING_ORACLE_BLOCKSIZE = 16

type PaddingOracle struct {
	Hostname   string `json:"hostname"`
	Port       int    `json:"port"`
	Iv         string `json:"iv"`
	Ciphertext string `json:"ciphertext"`
	Result     string `json:"result"`
}

func sendMessage(conn net.Conn, message []byte) error {
	_, err := conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}
	return nil
}

func receiveMessage(conn net.Conn, length int) ([]uint, error) {
	buffer := make([]byte, length)
	err := conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return []uint{}, fmt.Errorf("failed to set Timeout: %v", err)
	}
	_, err = conn.Read(buffer)
	if err != nil {
		return []uint{}, fmt.Errorf("failed to receive message: %v", err)
	}
	var res []uint
	for i, b := range buffer {
		if b == 1 {
			res = append(res, uint(i))
		}
	}
	return res, nil
}
func (p *PaddingOracle) Execute() {
	var plaintext []byte
	p.executeBlocks(&plaintext)
	p.Result = base64.StdEncoding.EncodeToString(plaintext)
}

func (p *PaddingOracle) executeBlocks(plaintext *[]byte) {
	initialCiphertext, _ := base64.StdEncoding.DecodeString(p.Ciphertext)

	iv, _ := base64.StdEncoding.DecodeString(p.Iv)

	for i := 0; i < len(initialCiphertext); i += PADDING_ORACLE_BLOCKSIZE {
		conn, err := net.Dial("tcp", p.Hostname+":"+fmt.Sprint(p.Port))
		if err != nil {
			panic(fmt.Sprintf("Error: failed to connect to server \n %v:", err))
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				panic(fmt.Sprintf("Error: failed to close connection \n %v:", err))
			}
		}(conn)
		ciphertext := initialCiphertext[i : i+PADDING_ORACLE_BLOCKSIZE]
		err = sendMessage(conn, ciphertext)
		if err != nil {
			panic(fmt.Sprintf("Error send init ciphertext message: %v", err))
		}
		plaintextBlock := make([]byte, PADDING_ORACLE_BLOCKSIZE)
		qBlocks := make([]byte, PADDING_ORACLE_BLOCKSIZE)

		p.executeByteIndex(conn, plaintextBlock, qBlocks)

		plaintextBlock = utils.Xor(plaintextBlock, iv)
		iv = ciphertext
		*plaintext = append(*plaintext, plaintextBlock...)
	}
}

func (p *PaddingOracle) executeByteIndex(conn net.Conn, plaintextBlock, qBlocks []byte) {
	for byteIndex := 1; byteIndex <= PADDING_ORACLE_BLOCKSIZE; byteIndex++ {
		var endI int
		var startI int
		endI = 256
		startI = 0
		// Step 2: Send length field (2 bytes, little endian)
		lengthField := make([]byte, 2)
		binary.LittleEndian.PutUint16(lengthField, uint16(endI-startI))
		err := sendMessage(conn, lengthField)
		if err != nil {
			panic(fmt.Sprintf("byteExecute: %v", err))
		}
		for i := startI; i < endI; i++ {
			// Step 3: Send Q-block (16 bytes)
			qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex] = byte(i)
			err = sendMessage(conn, qBlocks)
			if err != nil {
				panic(fmt.Sprintf("Error sending qBlock ByteIndex: %d, i: %d:\n %v", byteIndex, i, err))
			}
		}
		// Step 4: Receive response (l bytes)
		response, err := receiveMessage(conn, endI-startI)
		if err != nil {
			panic(fmt.Sprintf("Error: %v", err))
		}
		if len(response) == 0 {
			_, _ = fmt.Fprintf(os.Stderr, "startI %v \n", startI)
			_, _ = fmt.Fprintf(os.Stderr, "endI %v \n", endI)
			_, _ = fmt.Fprintf(os.Stderr, "ByteIndex %v \n", byteIndex)
			_, _ = fmt.Fprintf(os.Stderr, "QBlocks %v \n", qBlocks)
			_, _ = fmt.Fprintf(os.Stderr, "PlaintextBlock %v \n", plaintextBlock)
		}
		//multiple true responses
		if len(response) == 2 {
			newLengthField := make([]byte, 2)
			binary.LittleEndian.PutUint16(newLengthField, uint16(len(response)))
			err = sendMessage(conn, newLengthField)
			if err != nil {
				panic(fmt.Sprintf("multiple true check: Error sending qBlock ByteIndex: %d, i: %v:\n %v", byteIndex, response, err))
			}
			for _, q := range response {
				qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex-1] = byte(q ^ 0xff)
				qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex] = byte(q)
				err = sendMessage(conn, qBlocks)
				if err != nil {
					panic(fmt.Sprintf("multiple true check: Error sending qBlock ByteIndex: %d, i: %d:\n %v", byteIndex, q, err))
				}
			}
			NewResponse, err := receiveMessage(conn, len(response))
			if err != nil {
				panic(fmt.Sprintf("Error: %v", err))
			}

			response[0] = response[NewResponse[0]]
		}
		// berechnen von D(c)i = pi xor qi
		pByte := byte(byteIndex)
		q := byte(response[0])
		dc := pByte ^ q
		plaintextBlock[PADDING_ORACLE_BLOCKSIZE-byteIndex] = dc
		pNext := byte(byteIndex + 1)
		for j := 1; j <= byteIndex; j++ {
			qNext := pNext ^ plaintextBlock[PADDING_ORACLE_BLOCKSIZE-j]
			qBlocks[PADDING_ORACLE_BLOCKSIZE-j] = qNext
		}
	}
}
