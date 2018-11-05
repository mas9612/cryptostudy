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
			rc = mul(rc, 2)
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

func subBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = sbox[x][y]
	}
}

func shiftRows(state []byte) {
	l := len(state)
	t := make([]byte, l)
	copy(t, state)

	for y := 0; y < Nb; y++ {
		for x := 1; x < BytesOfWords; x++ {
			state[x+y*Nb] = t[x+Nb*((x+y)%BytesOfWords)]
		}
	}
}

func mixColumns(state []byte) {
	l := len(state)
	tmp := make([]byte, l)
	copy(tmp, state)

	for y := 0; y < Nb; y++ {
		for x := 0; x < l/Nb; x++ {
			state[y*BytesOfWords+x] = mul(polyMatrix[x][0], tmp[y*BytesOfWords]) ^ mul(polyMatrix[x][1], tmp[y*BytesOfWords+1]) ^ mul(polyMatrix[x][2], tmp[y*BytesOfWords+2]) ^ mul(polyMatrix[x][3], tmp[y*BytesOfWords+3])
		}
	}
}

func addRoundKey(state, key []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		state[i] ^= key[i]
	}
}

func invShiftRows(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for i := 1; i < BytesOfWords; i++ { // i is a row index
		colOffset := i
		for j := 0; j < Nb; j++ { // j is a column index
			column := (j + colOffset) % Nb
			state[column*4+i] = tmp[j*4+i]
		}
	}
}

func invSubBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = invSbox[x][y]
	}
}

func invMixColumns(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for i := 0; i < Nb; i++ {
		state[i*4] = mul(tmp[i*4], 0x0e) ^ mul(tmp[i*4+1], 0x0b) ^ mul(tmp[i*4+2], 0x0d) ^ mul(tmp[i*4+3], 0x09)
		state[i*4+1] = mul(tmp[i*4+1], 0x0e) ^ mul(tmp[i*4+2], 0x0b) ^ mul(tmp[i*4+3], 0x0d) ^ mul(tmp[i*4], 0x09)
		state[i*4+2] = mul(tmp[i*4+2], 0x0e) ^ mul(tmp[i*4+3], 0x0b) ^ mul(tmp[i*4], 0x0d) ^ mul(tmp[i*4+1], 0x09)
		state[i*4+3] = mul(tmp[i*4+3], 0x0e) ^ mul(tmp[i*4], 0x0b) ^ mul(tmp[i*4+1], 0x0d) ^ mul(tmp[i*4+2], 0x09)
	}
}
