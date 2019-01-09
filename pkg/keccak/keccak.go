package keccak

// Word is 64 bits in SHA-3
type Word [8]byte

// Lane size is Word size
type Lane Word

// State is a two-dimensional array of lanes
type State []byte

// Keccak calculates the hash value of sha3
func Keccak(d int, M []byte) ([]byte, error) {
	return nil, nil
}

// Theta is to XOR each bit in the state with the parities of two columns in the array.
func Theta(input State) (State, error) {

	return make(State, 0), nil
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
