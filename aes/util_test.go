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

func TestXtime(t *testing.T) {
	inputs := []byte{
		0x01,
		0x02,
		0x04,
		0x08,
		0x10,
		0x20,
		0x40,
		0x80,
		0x1b,
		0x36,
		0x6c,
		0xd8,
		0xab,
		0x4d,
		0x9a,
	}

	expected := []byte{
		0x02,
		0x04,
		0x08,
		0x10,
		0x20,
		0x40,
		0x80,
		0x1b,
		0x36,
		0x6c,
		0xd8,
		0xab,
		0x4d,
		0x9a,
		0x2f,
	}

	for i, input := range inputs {
		result := xtime(input)
		if result != expected[i] {
			t.Errorf("[TestXtime] case %d failed: result != expected : '%v' != '%v'", i, result, expected[i])
		}
	}
}
