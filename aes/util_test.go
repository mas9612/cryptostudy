package aes

import "testing"

func TestMul(t *testing.T) {
	inputs := [][]byte{
		[]byte{0x57, 0x01},
		[]byte{0x57, 0x02},
		[]byte{0x57, 0x04},
		[]byte{0x57, 0x08},
		[]byte{0x57, 0x10},
		[]byte{0x57, 0x13},
		[]byte{0x57, 0x83},
	}
	expected := []byte{
		0x57,
		0xae,
		0x47,
		0x8e,
		0x07,
		0xfe,
		0xc1,
	}

	for i, input := range inputs {
		result := mul(input[0], input[1])
		if result != expected[i] {
			t.Errorf("[TestMul] case %d failed: result != expected : '%v' != '%v'", i, result, expected[i])
		}
	}
}
