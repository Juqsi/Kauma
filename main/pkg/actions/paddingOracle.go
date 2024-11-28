package actions

import (
	"Abgabe/main/pkg/utils"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net"
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
	n, err := conn.Read(buffer)
	if err != nil {
		return []uint{}, fmt.Errorf("failed to receive message: %v", err)
	}
	if n != length {
		return []uint{}, fmt.Errorf("received message length mismatch: expected %d, got %d", length, n)
	}
	var res []uint
	for i, b := range buffer {
		if b&1 == 1 {
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

		plaintextBlock := p.executeByteIndex(conn)

		plaintextBlock = utils.Xor(plaintextBlock, iv)
		iv = ciphertext
		*plaintext = append(*plaintext, plaintextBlock...)
	}
}

func (p *PaddingOracle) executeByteIndex(conn net.Conn) []byte {
	plaintextBlock := make([]byte, PADDING_ORACLE_BLOCKSIZE)
	qBlocks := make([]byte, PADDING_ORACLE_BLOCKSIZE)

	for byteIndex := 1; byteIndex <= PADDING_ORACLE_BLOCKSIZE; byteIndex++ {
		p.processByteIndex(conn, byteIndex, plaintextBlock, qBlocks)
	}
	return plaintextBlock
}

func (p *PaddingOracle) processByteIndex(conn net.Conn, byteIndex int, plaintextBlock, qBlocks []byte) {
	amount := 32
	for i := 0; i < 256/amount; i++ {
		startI, endI := i*amount, (i+1)*amount
		p.sendLengthField(conn, amount)
		p.sendQBlocks(conn, byteIndex, qBlocks, startI, endI)
		response := p.receiveResponse(conn, amount)
		if len(response) < 1 {
			continue
		}
		if byteIndex == 1 {
			response = p.handleMultipleTrue(conn, byteIndex, qBlocks, response, startI)
		}
		p.calculatePlaintextBlock(byteIndex, response, startI, plaintextBlock, qBlocks)
		break
	}
}

func (p *PaddingOracle) sendLengthField(conn net.Conn, amount int) {
	lengthField := make([]byte, 2)
	binary.LittleEndian.PutUint16(lengthField, uint16(amount))
	err := sendMessage(conn, lengthField)
	if err != nil {
		panic(fmt.Sprintf("byteExecute: %v", err))
	}
}

func (p *PaddingOracle) sendQBlocks(conn net.Conn, byteIndex int, qBlocks []byte, startI, endI int) {
	for j := startI; j < endI; j++ {
		qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex] = byte(j)
		err := sendMessage(conn, qBlocks)
		if err != nil {
			panic(fmt.Sprintf("Error sending qBlock ByteIndex: %d, i: %d:\n %v", byteIndex, j, err))
		}
	}
}

func (p *PaddingOracle) receiveResponse(conn net.Conn, amount int) []uint {
	response, err := receiveMessage(conn, amount)
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}
	return response
}

func (p *PaddingOracle) handleMultipleTrue(conn net.Conn, byteIndex int, qBlocks []byte, response []uint, startI int) []uint {
	newLengthField := make([]byte, 2)
	binary.LittleEndian.PutUint16(newLengthField, uint16(len(response)))
	err := sendMessage(conn, newLengthField)
	if err != nil {
		panic(fmt.Sprintf("multiple true check: Error sending qBlock ByteIndex: %d, i: %v:\n %v", byteIndex, response, err))
	}
	for _, q := range response {
		q += uint(startI)
		qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex-1] = byte(q ^ 0xff)
		qBlocks[PADDING_ORACLE_BLOCKSIZE-byteIndex] = byte(q)
		err = sendMessage(conn, qBlocks)
		if err != nil {
			panic(fmt.Sprintf("multiple true check: Error sending qBlock ByteIndex: %d, i: %d:\n %v", byteIndex, q, err))
		}
	}
	newResponse, err := receiveMessage(conn, len(response))
	if err != nil {
		panic(fmt.Sprintf("Error: %v", err))
	}
	if len(newResponse) != 1 {
		return response
	}
	response[0] = response[newResponse[0]]
	return response
}

func (p *PaddingOracle) calculatePlaintextBlock(byteIndex int, response []uint, startI int, plaintextBlock, qBlocks []byte) {
	pByte := byte(byteIndex)
	q := byte(response[0] + uint(startI))
	dc := pByte ^ q
	plaintextBlock[PADDING_ORACLE_BLOCKSIZE-byteIndex] = dc
	pNext := byte(byteIndex + 1)
	for j := 1; j <= byteIndex; j++ {
		qNext := pNext ^ plaintextBlock[PADDING_ORACLE_BLOCKSIZE-j]
		qBlocks[PADDING_ORACLE_BLOCKSIZE-j] = qNext
	}
}
