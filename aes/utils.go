package aes

func mul(num1, num2 byte) byte {
	switch num2 {
	case 0:
		return 0
	case 1:
		return num1
	case 2:
		return mul2(num1)
	}

	remains := int(num2)
	result := num1
	i := 2
	for ; i <= remains; i *= 2 {
		result = mul2(result)
	}
	remains -= i / 2
	if remains > 0 {
		result ^= mul(num1, byte(remains))
	}
	return result
}

func mul2(num byte) byte {
	tmp := int(num) << 1 // multiplied by 2
	if tmp&0x100 != 0 {  //if 0x100 is set
		tmp ^= 0x11b // mod by 0x11b (x^8 + x^4 + x^3 + x + 1)
	}
	return byte(tmp)
}
