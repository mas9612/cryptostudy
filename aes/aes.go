package aes

import (
	"fmt"
	"log"
)

// Cipher encrypts plain text
func Cipher(in []byte, key []byte, mode int, iv []byte) []byte {
	switch len(key) {
	case 16:
		Nk = KeyLength128
		Nb = BlockSize128
		Nr = NumOfRounds128
	case 24:
		Nk = KeyLength192
		Nb = BlockSize192
		Nr = NumOfRounds192
	case 32:
		Nk = KeyLength256
		Nb = BlockSize256
		Nr = NumOfRounds256
	default:
		log.Fatalln("AES key length must be one of 128, 192, 256 bit")
	}

	expandedKey := make([]byte, BytesOfWords*Nb*(Nr+1))
	keyExpansion(key, expandedKey)

	numOfBlocks := len(in) / (Nb * BytesOfWords)
	if len(in)%(Nb*BytesOfWords) != 0 {
		numOfBlocks++
	}

	var out []byte
	switch mode {
	case ModeECB:
		out = ecbCipher(in, expandedKey, numOfBlocks)
	case ModeCBC:
		out = cbcCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCFB:
		out = cfbCipher(in, expandedKey, iv, numOfBlocks)
	case ModeOFB:
		out = ofbCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCTR:
		out = ctrCipher(in, expandedKey, iv, numOfBlocks)
	default:
		log.Fatalln("Invalid encryption mode")
	}
	return out
}

// InvCipher decrypt given cipher text
func InvCipher(in, key []byte, mode int, iv []byte) []byte {
	switch len(key) {
	case 16:
		Nk = KeyLength128
		Nb = BlockSize128
		Nr = NumOfRounds128
	case 24:
		Nk = KeyLength192
		Nb = BlockSize192
		Nr = NumOfRounds192
	case 32:
		Nk = KeyLength256
		Nb = BlockSize256
		Nr = NumOfRounds256
	default:
		log.Fatalln("AES key length must be one of 128, 192, 256 bit")
	}

	expandedKey := make([]byte, BytesOfWords*Nb*(Nr+1))
	keyExpansion(key, expandedKey)

	numOfBlocks := len(in) / (Nb * BytesOfWords)
	if len(in)%(Nb*BytesOfWords) != 0 {
		numOfBlocks++
	}

	var out []byte
	switch mode {
	case ModeECB:
		out = ecbInvCipher(in, expandedKey, numOfBlocks)
	case ModeCBC:
		out = cbcInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCFB:
		out = cfbInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeOFB:
		out = ofbInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCTR:
		out = ctrInvCipher(in, expandedKey, iv, numOfBlocks)
	default:
		log.Fatalln("Invalid encryption mode")
	}
	return out
}

func blockCipher(state, key []byte) {
	round := 0
	if round == PrintNRound {
		fmt.Printf("[Round %d]\n", round)
	}
	addRoundKey(state, key[:Nb*BytesOfWords])
	if round == PrintNRound {
		fmt.Print("After AddRoundKey: ")
		printBytes(state)
	}

	for round = 1; round <= Nr; round++ {
		if round == PrintNRound {
			fmt.Printf("[Round %d]\n", round)
		}
		subBytes(state)
		if round == PrintNRound {
			fmt.Print("After SubBytes: ")
			printBytes(state)
		}
		shiftRows(state)
		if round == PrintNRound {
			fmt.Print("After ShiftRows: ")
			printBytes(state)
		}
		if round < Nr {
			mixColumns(state)
			if round == PrintNRound {
				fmt.Print("After MixColumns: ")
				printBytes(state)
			}
		}
		addRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
		if round == PrintNRound {
			fmt.Print("After AddRoundKey: ")
			printBytes(state)
		}
	}
}

func invCipher(state, key []byte) {
	round := Nr
	if round == PrintNRound {
		fmt.Printf("[Round %d]\n", round)
	}
	addRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
	if round == PrintNRound {
		fmt.Print("After AddRoundKey: ")
		printBytes(state)
	}

	for round = Nr - 1; round >= 0; round-- {
		if round == PrintNRound {
			fmt.Printf("[Round %d]\n", round)
		}
		invShiftRows(state)
		if round == PrintNRound {
			fmt.Print("After InvShiftRows: ")
			printBytes(state)
		}
		invSubBytes(state)
		if round == PrintNRound {
			fmt.Print("After InvSubBytes: ")
			printBytes(state)
		}
		addRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
		if round == PrintNRound {
			fmt.Print("After AddRoundKey: ")
			printBytes(state)
		}
		if round > 0 {
			invMixColumns(state)
			if round == PrintNRound {
				fmt.Print("After InvMixColumns: ")
				printBytes(state)
			}
		}
	}
}
