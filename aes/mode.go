package aes

import (
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
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		invCipher(state, key)
		// XOR with previous cipher block
		for j := 0; j < Nb*BytesOfWords; j++ {
			state[j] ^= previous[j]
		}
		copy(out[from:from+Nb*BytesOfWords], state)
		copy(previous, state)
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

func cfbCipher(in, key, feedback []byte) []byte {
	return in
}

func ofbCipher(in, key, feedback []byte) []byte {
	return in
}

func ctrCipher(in, key, feedback []byte) []byte {
	return in
}
