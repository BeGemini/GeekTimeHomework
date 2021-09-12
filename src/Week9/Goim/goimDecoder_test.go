package Goim

import (
	"encoding/binary"
	"testing"
)

func encoder(msg string) []byte {
	headerLen := 16
	version := 1
	operation := 3
	sequence := 10086
	packageLen := headerLen + len(msg)
	pack := make([]byte, packageLen)
	binary.BigEndian.PutUint32(pack[:4], uint32(packageLen))
	binary.BigEndian.PutUint16(pack[4:6], uint16(headerLen))
	binary.BigEndian.PutUint16(pack[6:8], uint16(version))
	binary.BigEndian.PutUint32(pack[8:12], uint32(operation))
	binary.BigEndian.PutUint32(pack[12:16], uint32(sequence))
	copy(pack[16:], []byte(msg))
	return pack
}

func TestDecoder(t *testing.T) {
	msg := ""
	decoder(encoder(msg))
}
