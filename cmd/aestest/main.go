package main

import (
	"fmt"

	"github.com/mas9612/cryptostudy/aes"
)

func main() {
	plainText := []byte{0x00, 0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90, 0xa0, 0xb0, 0xc0, 0xd0, 0xe0, 0xf0}
	fmt.Println("Before Cipher:")
	for _, b := range plainText {
		fmt.Printf("%02x ", b)
	}
	fmt.Print("\n")

	cipherText := make([]byte, len(plainText))
	aes.Cipher(plainText, cipherText, []byte{})
	fmt.Println("After Cipher:")
	for _, b := range cipherText {
		fmt.Printf("%02x ", b)
	}
	fmt.Print("\n")
}
