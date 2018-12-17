package aes

import (
	"fmt"
)

// Xtime calculate multiplis n and 2 in a Galois Field
func Xtime(n byte) byte {
	p := int(n) << 1
	if p&0x100 != 0 {
		p ^= poly
	}
	return byte(p)
}

// Mul multiplies n and p in a Galois Field
func Mul(n, p byte) byte {
	switch {
	case p == 0:
		return 0
	case p%2 == 0:
		t := Mul(n, p/2)
		return Xtime(t)
	case p%2 == 1:
		return n ^ Mul(n, p-1)
	}
	return n
}

// PrintableBytes returns printable string from []byte
func PrintableBytes(bytes []byte) (str string) {
	str = ""
	for i, b := range bytes {
		str += fmt.Sprintf("%#02x", b)
		if i != len(bytes)-1 {
			str += " "
		}
	}
	return
}

func printRoundBytes(bytes []byte, round int, phase string) {
	if round == PrintNRound {
		fmt.Printf("After %s: %s\n", phase, PrintableBytes(bytes))
	}
}
