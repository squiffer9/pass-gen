package utils

import (
	"encoding/binary"
	"encoding/hex"
)

func HexToInt64(hexStr string) int64 {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	// Ensure we have exactly 8 bytes
	if len(bytes) != 8 {
		panic("Invalid hex string length")
	}

	return int64(binary.BigEndian.Uint64(bytes))
}
