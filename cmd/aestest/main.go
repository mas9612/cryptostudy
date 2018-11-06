package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/mas9612/cryptostudy/aes"
)

func main() {
	fs := flag.NewFlagSet("AES", flag.ExitOnError)
	key := fs.String("K", "", "Encrypt key (hexadecimal notation)")
	mode := fs.String("mode", "", "Cipher mode. Valid mode is one of [ECB, CBC, CFB, OFB, CTR]")
	iv := fs.String("iv", "", "IV")
	round := fs.Int("r", -1, "Print round N result")
	decrypt := fs.Bool("d", false, "Decrypt (Default Encrypt)")
	help := fs.Bool("help", false, "Print help and exit")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Failed to parse command line arguments")
		os.Exit(1)
	}

	if *help {
		fs.Usage()
		os.Exit(0)
	}

	if *key == "" {
		fmt.Println("Missing -K")
		os.Exit(1)
	} else if len(*key) != 32 && len(*key) != 48 && len(*key) != 64 {
		fmt.Println("Key must be one of 16, 24, 32 bytes length: ", len(*key)/2)
		os.Exit(1)
	}

	if *mode == "" {
		fmt.Println("Missing -mode")
		os.Exit(1)
	}
	switch *mode {
	case "CBC", "CFB", "OFB", "CTR":
		if *iv == "" {
			fmt.Println("Missing -iv")
			os.Exit(1)
		} else if len(*iv) != 32 {
			fmt.Println("IV must be 16 bytes length")
			os.Exit(1)
		}
	case "ECB":
	default:
		fmt.Println("Invalid mode")
		os.Exit(1)
	}

	var cipherMode int
	switch *mode {
	case "ECB":
		cipherMode = aes.ModeECB
	case "CBC":
		cipherMode = aes.ModeCBC
	case "CFB":
		cipherMode = aes.ModeCFB
	case "OFB":
		cipherMode = aes.ModeOFB
	case "CTR":
		cipherMode = aes.ModeCTR
	}

	aes.PrintNRound = *round

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Failed to read from stdin")
		os.Exit(1)
	}

	var result []byte
	if !*decrypt {
		result = aes.Cipher(bytes, hexStringToBytes(*key), cipherMode, hexStringToBytes(*iv))
	} else {
		result = aes.InvCipher(bytes, hexStringToBytes(*key), cipherMode, hexStringToBytes(*iv))
	}
	fmt.Print(string(result))
}

func hexStringToBytes(hex string) []byte {
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
