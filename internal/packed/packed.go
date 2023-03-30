package packed

import (
	"bytes"
	"encoding/binary"
)

// Pack 将数据打包成二进制格式。
func Pack(v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unpack 从二进制格式解包数据。
func Unpack(data []byte, v interface{}) error {
	buf := bytes.NewBuffer(data)
	return binary.Read(buf, binary.LittleEndian, v)
}
