package aes

import (
	"fmt"
)

func xtime(n byte) byte {
	p := int(n) << 1
	if p&0x100 != 0 {
		p ^= poly
	}
	return byte(p)
}

func mul(n, p byte) byte {
	switch {
	case p == 0:
		return 0
	case p%2 == 0:
		t := mul(n, p/2)
		return xtime(t)
	case p%2 == 1:
		return n ^ mul(n, p-1)
	}
	return n
}

func printableBytes(bytes []byte) (str string) {
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
		fmt.Printf("After %s: %s\n", phase, printableBytes(bytes))
	}
}
