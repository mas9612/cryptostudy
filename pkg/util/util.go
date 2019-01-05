package util

import (
	"fmt"
	"os"
	"strconv"
)

// HexStringToBytes returns slice of bytes converted from given hex string
func HexStringToBytes(hex string) []byte {
	bytes := make([]byte, len(hex)/2)
	idx := 0

	for i := 0; i < len(hex); i += 2 {
		tmp, err := strconv.ParseUint(hex[i:i+2], 16, 8)
		if err != nil {
			fmt.Println("Failed to parse hex string")
			os.Exit(1)
		}
		bytes[idx] = byte(tmp)
		idx++
	}
	return bytes
}
