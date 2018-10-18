package main

import (
	"fmt"

	"github.com/mas9612/cryptostudy/aes"
)

func main() {
	plain := "This is a test!!"
	bytePlain := []byte(plain)
	fmt.Printf("Before Cipher: %v, length: %v\n", bytePlain, len(bytePlain))

	aes.Cipher(bytePlain)
	fmt.Printf("After Cipher: %v, length: %v\n", bytePlain, len(bytePlain))
}
