package aes

import (
	"encoding/binary"
	"log"
)

func ecbCipher(in, key []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)

	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		// add padding if need
		if stateLength < Nb*BytesOfWords {
			padding := Nb*BytesOfWords - stateLength
			for i := stateLength; i < Nb*BytesOfWords; i++ {
				state[i] = byte(padding)
			}
		}

		cipher(state, key)
		copy(out[from:from+Nb*BytesOfWords], state)
	}
	return out
}

func ecbInvCipher(in, key []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)

	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		invCipher(state, key)
		copy(out[from:from+Nb*BytesOfWords], state)
	}

	// determine last byte (to remove padding)
	var end int
	if int(out[len(out)-1]) < Nb*BytesOfWords {
		padding := int(out[len(out)-1])
		end = len(out) - padding
	} else {
		end = len(out)
	}
	return out[:end]
}

func cbcCipher(in, key, iv []byte, numOfBlocks int) []byte {
	if len(iv) != Nb*BytesOfWords {
		log.Fatalf("IV must be same as block size (%d byte)", Nb*BytesOfWords)
	}

	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)
	copy(previous, iv)

	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		// add padding if need
		if stateLength < Nb*BytesOfWords {
			padding := Nb*BytesOfWords - stateLength
			for i := stateLength; i < Nb*BytesOfWords; i++ {
				state[i] = byte(padding)
			}
		}

		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}

		cipher(state, key)
		copy(out[from:from+Nb*BytesOfWords], state)
		// save cipher block for next block
		copy(previous, state)
	}
	return out
}

func cbcInvCipher(in, key, iv []byte, numOfBlocks int) []byte {
	if len(iv) != Nb*BytesOfWords {
		log.Fatalf("IV must be same as block size (%d byte)", Nb*BytesOfWords)
	}

	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)
	copy(previous, iv)

	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		tmp := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])
		copy(tmp, state)

		invCipher(state, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)
		copy(previous, tmp)
	}

	// determine last byte (to remove padding)
	var end int
	if int(out[len(out)-1]) < Nb*BytesOfWords {
		padding := int(out[len(out)-1])
		end = len(out) - padding
	} else {
		end = len(out)
	}
	return out[:end]
}

func cfbCipher(in, key, iv []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)
	copy(previous, iv)

	var end int
	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		cipher(previous, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)
		copy(previous, state)

		if stateLength < Nb*BytesOfWords {
			end = Nb*BytesOfWords - stateLength
		}
	}
	return out[:len(out)-end]
}

func cfbInvCipher(in, key, iv []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)
	copy(previous, iv)

	var end int
	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		tmp := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])
		copy(tmp, state)

		cipher(previous, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)
		copy(previous, tmp)

		if stateLength < Nb*BytesOfWords {
			end = Nb*BytesOfWords - stateLength
		}
	}
	return out[:len(out)-end]
}

func ofbCipher(in, key, iv []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)
	copy(previous, iv)

	var end int
	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		cipher(previous, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)

		if stateLength < Nb*BytesOfWords {
			end = Nb*BytesOfWords - stateLength
		}
	}
	return out[:len(out)-end]
}

func ofbInvCipher(in, key, iv []byte, numOfBlocks int) []byte {
	return ofbCipher(in, key, iv, numOfBlocks)
}

func ctrCipher(in, key, iv []byte, numOfBlocks int) []byte {
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)
	previous := make([]byte, Nb*BytesOfWords)

	nonce := make([]byte, (Nb/2)*BytesOfWords)
	copy(nonce, iv[:(Nb/2)*BytesOfWords])
	counter := binary.BigEndian.Uint64(iv[(Nb/2)*BytesOfWords:])

	var end int
	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		copy(previous, nonce)
		binary.BigEndian.PutUint64(previous[(Nb/2)*BytesOfWords:], uint64(counter))
		cipher(previous, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)
		counter++

		if stateLength < Nb*BytesOfWords {
			end = Nb*BytesOfWords - stateLength
		}
	}
	return out[:len(out)-end]
}

func ctrInvCipher(in, key, iv []byte, numOfBlocks int) []byte {
	return ctrCipher(in, key, iv, numOfBlocks)
}
