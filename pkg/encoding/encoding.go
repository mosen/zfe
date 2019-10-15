package encoding

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"go/types"
)

var HEADER = []byte{'Z', 'B', 'X', 'D', 0x01}

func Encode(data interface{}) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	bodyLength := uint64(len(body))
	buf := bytes.NewBuffer(HEADER)
	var bodyLengthBytes = make([]byte, binary.Size(types.Uint64))
	binary.LittleEndian.PutUint64(bodyLengthBytes, bodyLength)
	buf.Write(bodyLengthBytes)
	buf.Write(body)

	return buf.Bytes(), nil
}
