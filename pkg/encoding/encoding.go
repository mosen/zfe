package encoding

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

var HEADER = []byte{'Z', 'B', 'X', 'D', 0x01}

func Encode(data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	bodyLength := uint64(len(body))
	fmt.Println(bodyLength)
	buf := bytes.NewBuffer(HEADER)

	lengthBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(lengthBytes, bodyLength)

	buf.Write(lengthBytes)
	buf.Write(body)

	return buf.Bytes(), nil
}

func DecodeNext(reader *bufio.Reader) ([]byte, error) {
	headerAndLength := make([]byte, 8+len(HEADER))
	if _, err := reader.Read(headerAndLength); err != nil {
		return nil, err
	}

	fmt.Println(headerAndLength)
	return nil, nil
}
