package keccak

import (
	"fmt"
)

// Word is 64 bits in SHA-3
type Word [8]byte

// Lane size is Word size
type Lane Word

// Plane is x-z plane of State
type Plane [5]Lane

func (p *Plane) get(x, z int) byte {
	// center of x-axis is labeled as index 0
	x = (x + 2) % 5
	byteIndex := z / 8
	bitIndex := z % 8
	return (p[x][byteIndex] >> uint(7-bitIndex)) & 0x1
}

func (p *Plane) set(x, z int, bit byte) error {
	// center of x-axis is labeled as index 0
	x = (x + 2) % 5
	byteIndex := z / 8
	bitIndex := z % 8
	if bit == 0x1 {
		p[x][byteIndex] |= 0x1
	} else if bit == 0x0 {
		mask := 0x1 << uint(7-bitIndex)
		p[x][byteIndex] &= byte(mask ^ 0xff)
	} else {
		return fmt.Errorf("The value passed to the set of Plane is neither 0 nor 1. Set Plane to 0 or 1 of type byte")
	}
	return nil
}

// State is a two-dimensional array of lanes
type State []byte

func (s *State) get(x, y, z int) byte {
	bitIndex := w*(5*y+x) + z
	byteIndex := bitIndex / 8
	index := 7 - (bitIndex % 8)
	b := (*s)[byteIndex]
	mask := byte(1 << uint(index))

	return (b & mask) >> uint(index)
}
func (s *State) set(x, y, z int, bit byte) error {
	bitIndex := w*(5*y+x) + z
	byteIndex := bitIndex / 8
	index := 7 - (bitIndex % 8)
	b := (*s)[byteIndex]
	mask := byte(1 << uint(index))
	if bit == byte(1) {
		(*s)[byteIndex] = b | mask
	} else if bit == byte(0) {
		(*s)[byteIndex] = b & (mask ^ byte(0xff))
	} else {
		return fmt.Errorf("The value passed to the set of State is neither 0 nor 1.Set State to 0 or 1 of type byte")
	}
	return nil
}

// Keccak calculates the hash value of sha3
func Keccak(d int, M []byte) ([]byte, error) {
	return nil, nil
}

// Theta is to XOR each bit in the state with the parities of two columns in the array.
func Theta(input State) (State, error) {
	var C Plane
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			var bit byte
			for i := 0; i < 5; i++ {
				bit ^= input.get(x, i, z)
			}
			C.set(x, z, bit)
		}
	}
	var D Plane
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			a := C.get(Modulo(x-1, 5), z)
			b := C.get(Modulo(x+1, 5), Modulo(z-1, w))
			D.set(x, z, a^b)
		}
	}
	state := make(State, stateByteLen)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				bit := input.get(x, y, z) ^ D.get(x, z)
				state.set(x, y, z, bit)
			}
		}
	}
	return state, nil
}

// Rho is to rotate the bits of each lane by a length, called the offset, which depends on the fixed x and y coordinates of the lane.
func Rho(input State) (State, error) {

	return make(State, 0), nil
}

// Pi is to rearrange the positions of the lanes in any slice.
func Pi(input State) (State, error) {

	return make(State, 0), nil
}

// Chi is to XOR each bit with a non-linear function of two other bits in its row.
func Chi(input State) (State, error) {

	return make(State, 0), nil
}

// Iota is to modify some of the bits of Lane (0, 0) in a manner that depends on the round index i_r. The other 24 lanes are not affected by Iota.
func Iota(input State) (State, error) {

	return make(State, 0), nil
}

// Rc calculates a round constant.it use in Iota.
func Rc(t int) (byte, error) {

	return 0, nil
}

// Pad makes padding represented by a regular expression of 10*1.
func Pad(x, m int) ([]byte, error) {

	return make([]byte, 0), nil
}

// Modulo returns a mod b
func Modulo(a, b int) int {
	tmp := a % b
	if tmp < 0 {
		return tmp + b
	}
	return tmp
}
