package aes

func keyExpansion(key []byte, expanded []byte) {
	copy(expanded, key)

	rc := byte(1) // round constant

	for i := Nk; i < Nb*(Nr+1); i++ {
		tmp := make([]byte, 4)
		copy(tmp, expanded[i*4-BytesOfWords:i*4]) // copy previous word from expanded key to tmp
		if i%Nk == 0 {
			rotWord(tmp)
			subWord(tmp)
			tmp[0] ^= rc
			rc = Mul(rc, 2)
		} else if Nk > 6 && i%Nk == 4 {
			subWord(tmp)
		}

		for j := 0; j < BytesOfWords; j++ {
			expanded[i*4+j] = expanded[(i-Nk)*4+j] ^ tmp[j]
		}
	}
}

func rotWord(word []byte) {
	tmp := word[0]
	for i := 0; i < BytesOfWords-1; i++ {
		word[i] = word[i+1]
	}
	word[BytesOfWords-1] = tmp
}

func subWord(word []byte) {
	for i := 0; i < BytesOfWords; i++ {
		x := word[i] >> 4
		y := word[i] & 0xf
		word[i] = sbox[x][y]
	}
}

// SubBytes transforms given state with sbox
func SubBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = sbox[x][y]
	}
}

// ShiftRows transforms given state with byte shift
func ShiftRows(state []byte) {
	l := len(state)
	t := make([]byte, l)
	copy(t, state)

	for y := 0; y < Nb; y++ {
		for x := 1; x < BytesOfWords; x++ {
			state[x+y*Nb] = t[x+Nb*((x+y)%BytesOfWords)]
		}
	}
}

// MixColumns transforms given state with multiplication in a Golois Field
func MixColumns(state []byte) {
	l := len(state)
	tmp := make([]byte, l)
	copy(tmp, state)

	for y := 0; y < Nb; y++ {
		for x := 0; x < l/Nb; x++ {
			state[y*BytesOfWords+x] = Mul(polyMatrix[x][0], tmp[y*BytesOfWords]) ^ Mul(polyMatrix[x][1], tmp[y*BytesOfWords+1]) ^ Mul(polyMatrix[x][2], tmp[y*BytesOfWords+2]) ^ Mul(polyMatrix[x][3], tmp[y*BytesOfWords+3])
		}
	}
}

// AddRoundKey transforms given state with XOR to round key
func AddRoundKey(state, key []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		state[i] ^= key[i]
	}
}

// InvShiftRows transforms given state with byte shift
func InvShiftRows(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for y := 0; y < Nb; y++ {
		for x := 1; x < BytesOfWords; x++ {
			state[x+y*BytesOfWords] = tmp[x+BytesOfWords*((y+Nb-x)%BytesOfWords)]
		}
	}
}

// InvSubBytes transforms given state with sbox
func InvSubBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = invSbox[x][y]
	}
}

// InvMixColumns transforms given state with multiplication in a Golois Field
func InvMixColumns(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for y := 0; y < Nb; y++ {
		for x := 0; x < BytesOfWords; x++ {
			state[y*BytesOfWords+x] = Mul(invPolyMatrix[x][0], tmp[y*BytesOfWords]) ^ Mul(invPolyMatrix[x][1], tmp[y*BytesOfWords+1]) ^ Mul(invPolyMatrix[x][2], tmp[y*BytesOfWords+2]) ^ Mul(invPolyMatrix[x][3], tmp[y*BytesOfWords+3])
		}
	}
}
