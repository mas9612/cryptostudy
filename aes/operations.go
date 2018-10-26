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
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for i := 1; i < BytesOfWords; i++ { // i is a row index
		colOffset := i
		for j := 0; j < Nb; j++ { // j is a column index
			column := (j + colOffset) % Nb
			state[j*4+i] = tmp[column*4+i]
		}
	}
}

func mixColumns(state []byte) {
	bytes := make([]byte, Nb*BytesOfWords)
	copy(bytes, state)

	for i := 0; i < Nb; i++ {
		mulBy2 := make([]byte, BytesOfWords)
		for j := 0; j < BytesOfWords; j++ {
			mulBy2[j] = mul(bytes[i*4+j], 2)
		}

		state[i*4] = mulBy2[0] ^ (mulBy2[1] ^ bytes[i*4+1]) ^ bytes[i*4+2] ^ bytes[i*4+3]
		state[i*4+1] = mulBy2[1] ^ (mulBy2[2] ^ bytes[i*4+2]) ^ bytes[i*4+0] ^ bytes[i*4+3]
		state[i*4+2] = mulBy2[2] ^ (mulBy2[3] ^ bytes[i*4+3]) ^ bytes[i*4+0] ^ bytes[i*4+1]
		state[i*4+3] = mulBy2[3] ^ (mulBy2[0] ^ bytes[i*4+0]) ^ bytes[i*4+1] ^ bytes[i*4+2]
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
	bytes := make([]byte, Nb*BytesOfWords)
	copy(bytes, state)

	for i := 0; i < Nb; i++ {
		state[i*4] = mul(bytes[i*4], 0x0e) ^ mul(bytes[i*4+1], 0x0b) ^ mul(bytes[i*4+2], 0x0d) ^ mul(bytes[i*4+3], 0x09)
		state[i*4+1] = mul(bytes[i*4+1], 0x0e) ^ mul(bytes[i*4+2], 0x0b) ^ mul(bytes[i*4+3], 0x0d) ^ mul(bytes[i*4], 0x09)
		state[i*4+2] = mul(bytes[i*4+2], 0x0e) ^ mul(bytes[i*4+3], 0x0b) ^ mul(bytes[i*4], 0x0d) ^ mul(bytes[i*4+1], 0x09)
		state[i*4+3] = mul(bytes[i*4+3], 0x0e) ^ mul(bytes[i*4], 0x0b) ^ mul(bytes[i*4+1], 0x0d) ^ mul(bytes[i*4+2], 0x09)
	}
}