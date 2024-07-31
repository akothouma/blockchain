package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	if err := binary.Write(buff, binary.BigEndian, num); err != nil {
		fmt.Fprintln(os.Stderr,err)
	}
	return buff.Bytes()
}
