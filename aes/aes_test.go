package aes

import (
	"bytes"
	"testing"
)

func TestCipher(t *testing.T) {
	Nb = BlockSize128
	Nr = NumOfRounds128
	Nk = KeyLength128

	inputs := [][]byte{
		[]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34},
		[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
		[]byte{
			0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
			0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
		},
		[]byte{
			0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
			0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc,
		},
	}
	keys := [][]byte{
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
	}
	expected := [][]byte{
		[]byte{0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32},
		[]byte{0x69, 0xc4, 0xe0, 0xd8, 0x6a, 0x7b, 0x04, 0x30, 0xd8, 0xcd, 0xb7, 0x80, 0x70, 0xb4, 0xc5, 0x5a},
		[]byte{
			0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
			0x8d, 0xf4, 0xe9, 0xaa, 0xc5, 0xc7, 0x57, 0x3a, 0x27, 0xd8, 0xd0, 0x55, 0xd6, 0xe4, 0xd6, 0x4b,
		},
		[]byte{
			0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
			0xc7, 0x62, 0x4e, 0x2a, 0x6a, 0x9d, 0xff, 0xb3, 0x0a, 0xfb, 0x09, 0x44, 0x62, 0x01, 0x4c, 0xf0,
		},
	}

	for i, input := range inputs {
		cipherText := Cipher(input, keys[i])
		if !bytes.Equal(cipherText, expected[i]) {
			t.Errorf("[TestCipher] case %d failed: cipherText != expected : '%v' != '%v'", i, cipherText, expected[i])
		}
	}
}

func TestInvCipher(t *testing.T) {
	Nb = BlockSize128
	Nr = NumOfRounds128
	Nk = KeyLength128

	inputs := [][]byte{
		[]byte{0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32},
		[]byte{0x69, 0xc4, 0xe0, 0xd8, 0x6a, 0x7b, 0x04, 0x30, 0xd8, 0xcd, 0xb7, 0x80, 0x70, 0xb4, 0xc5, 0x5a},
		[]byte{
			0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
			0x8d, 0xf4, 0xe9, 0xaa, 0xc5, 0xc7, 0x57, 0x3a, 0x27, 0xd8, 0xd0, 0x55, 0xd6, 0xe4, 0xd6, 0x4b,
		},
		[]byte{
			0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32,
			0xc7, 0x62, 0x4e, 0x2a, 0x6a, 0x9d, 0xff, 0xb3, 0x0a, 0xfb, 0x09, 0x44, 0x62, 0x01, 0x4c, 0xf0,
		},
	}
	keys := [][]byte{
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f},
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
	}
	expected := [][]byte{
		[]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34},
		[]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
		[]byte{
			0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
			0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff,
		},
		[]byte{
			0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
			0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0xbb, 0xcc,
		},
	}

	for i, input := range inputs {
		plainText := InvCipher(input, keys[i])
		if !bytes.Equal(plainText, expected[i]) {
			t.Errorf("[TestInvCipher] case %d failed: plainText != expected : '%v' != '%v'", i, plainText, expected[i])
		}
	}
}

func TestKeyExpansion(t *testing.T) {
	keyLength := []int{
		KeyLength128,
		KeyLength192,
		KeyLength256,
	}
	inputs := [][]byte{
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{
			0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
			0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
		},
		[]byte{
			0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
			0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
		},
	}
	expected := [][]byte{
		[]byte{
			0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c,
			0xa0, 0xfa, 0xfe, 0x17, 0x88, 0x54, 0x2c, 0xb1, 0x23, 0xa3, 0x39, 0x39, 0x2a, 0x6c, 0x76, 0x05,
			0xf2, 0xc2, 0x95, 0xf2, 0x7a, 0x96, 0xb9, 0x43, 0x59, 0x35, 0x80, 0x7a, 0x73, 0x59, 0xf6, 0x7f,
			0x3d, 0x80, 0x47, 0x7d, 0x47, 0x16, 0xfe, 0x3e, 0x1e, 0x23, 0x7e, 0x44, 0x6d, 0x7a, 0x88, 0x3b,
			0xef, 0x44, 0xa5, 0x41, 0xa8, 0x52, 0x5b, 0x7f, 0xb6, 0x71, 0x25, 0x3b, 0xdb, 0x0b, 0xad, 0x00,
			0xd4, 0xd1, 0xc6, 0xf8, 0x7c, 0x83, 0x9d, 0x87, 0xca, 0xf2, 0xb8, 0xbc, 0x11, 0xf9, 0x15, 0xbc,
			0x6d, 0x88, 0xa3, 0x7a, 0x11, 0x0b, 0x3e, 0xfd, 0xdb, 0xf9, 0x86, 0x41, 0xca, 0x00, 0x93, 0xfd,
			0x4e, 0x54, 0xf7, 0x0e, 0x5f, 0x5f, 0xc9, 0xf3, 0x84, 0xa6, 0x4f, 0xb2, 0x4e, 0xa6, 0xdc, 0x4f,
			0xea, 0xd2, 0x73, 0x21, 0xb5, 0x8d, 0xba, 0xd2, 0x31, 0x2b, 0xf5, 0x60, 0x7f, 0x8d, 0x29, 0x2f,
			0xac, 0x77, 0x66, 0xf3, 0x19, 0xfa, 0xdc, 0x21, 0x28, 0xd1, 0x29, 0x41, 0x57, 0x5c, 0x00, 0x6e,
			0xd0, 0x14, 0xf9, 0xa8, 0xc9, 0xee, 0x25, 0x89, 0xe1, 0x3f, 0x0c, 0xc8, 0xb6, 0x63, 0x0c, 0xa6,
		},
		[]byte{
			0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
			0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b, 0xfe, 0x0c, 0x91, 0xf7, 0x24, 0x02, 0xf5, 0xa5,
			0xec, 0x12, 0x06, 0x8e, 0x6c, 0x82, 0x7f, 0x6b, 0x0e, 0x7a, 0x95, 0xb9, 0x5c, 0x56, 0xfe, 0xc2,
			0x4d, 0xb7, 0xb4, 0xbd, 0x69, 0xb5, 0x41, 0x18, 0x85, 0xa7, 0x47, 0x96, 0xe9, 0x25, 0x38, 0xfd,
			0xe7, 0x5f, 0xad, 0x44, 0xbb, 0x09, 0x53, 0x86, 0x48, 0x5a, 0xf0, 0x57, 0x21, 0xef, 0xb1, 0x4f,
			0xa4, 0x48, 0xf6, 0xd9, 0x4d, 0x6d, 0xce, 0x24, 0xaa, 0x32, 0x63, 0x60, 0x11, 0x3b, 0x30, 0xe6,
			0xa2, 0x5e, 0x7e, 0xd5, 0x83, 0xb1, 0xcf, 0x9a, 0x27, 0xf9, 0x39, 0x43, 0x6a, 0x94, 0xf7, 0x67,
			0xc0, 0xa6, 0x94, 0x07, 0xd1, 0x9d, 0xa4, 0xe1, 0xec, 0x17, 0x86, 0xeb, 0x6f, 0xa6, 0x49, 0x71,
			0x48, 0x5f, 0x70, 0x32, 0x22, 0xcb, 0x87, 0x55, 0xe2, 0x6d, 0x13, 0x52, 0x33, 0xf0, 0xb7, 0xb3,
			0x40, 0xbe, 0xeb, 0x28, 0x2f, 0x18, 0xa2, 0x59, 0x67, 0x47, 0xd2, 0x6b, 0x45, 0x8c, 0x55, 0x3e,
			0xa7, 0xe1, 0x46, 0x6c, 0x94, 0x11, 0xf1, 0xdf, 0x82, 0x1f, 0x75, 0x0a, 0xad, 0x07, 0xd7, 0x53,
			0xca, 0x40, 0x05, 0x38, 0x8f, 0xcc, 0x50, 0x06, 0x28, 0x2d, 0x16, 0x6a, 0xbc, 0x3c, 0xe7, 0xb5,
			0xe9, 0x8b, 0xa0, 0x6f, 0x44, 0x8c, 0x77, 0x3c, 0x8e, 0xcc, 0x72, 0x04, 0x01, 0x00, 0x22, 0x02,
		},
		[]byte{
			0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
			0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
			0x9b, 0xa3, 0x54, 0x11, 0x8e, 0x69, 0x25, 0xaf, 0xa5, 0x1a, 0x8b, 0x5f, 0x20, 0x67, 0xfc, 0xde,
			0xa8, 0xb0, 0x9c, 0x1a, 0x93, 0xd1, 0x94, 0xcd, 0xbe, 0x49, 0x84, 0x6e, 0xb7, 0x5d, 0x5b, 0x9a,
			0xd5, 0x9a, 0xec, 0xb8, 0x5b, 0xf3, 0xc9, 0x17, 0xfe, 0xe9, 0x42, 0x48, 0xde, 0x8e, 0xbe, 0x96,
			0xb5, 0xa9, 0x32, 0x8a, 0x26, 0x78, 0xa6, 0x47, 0x98, 0x31, 0x22, 0x29, 0x2f, 0x6c, 0x79, 0xb3,
			0x81, 0x2c, 0x81, 0xad, 0xda, 0xdf, 0x48, 0xba, 0x24, 0x36, 0x0a, 0xf2, 0xfa, 0xb8, 0xb4, 0x64,
			0x98, 0xc5, 0xbf, 0xc9, 0xbe, 0xbd, 0x19, 0x8e, 0x26, 0x8c, 0x3b, 0xa7, 0x09, 0xe0, 0x42, 0x14,
			0x68, 0x00, 0x7b, 0xac, 0xb2, 0xdf, 0x33, 0x16, 0x96, 0xe9, 0x39, 0xe4, 0x6c, 0x51, 0x8d, 0x80,
			0xc8, 0x14, 0xe2, 0x04, 0x76, 0xa9, 0xfb, 0x8a, 0x50, 0x25, 0xc0, 0x2d, 0x59, 0xc5, 0x82, 0x39,
			0xde, 0x13, 0x69, 0x67, 0x6c, 0xcc, 0x5a, 0x71, 0xfa, 0x25, 0x63, 0x95, 0x96, 0x74, 0xee, 0x15,
			0x58, 0x86, 0xca, 0x5d, 0x2e, 0x2f, 0x31, 0xd7, 0x7e, 0x0a, 0xf1, 0xfa, 0x27, 0xcf, 0x73, 0xc3,
			0x74, 0x9c, 0x47, 0xab, 0x18, 0x50, 0x1d, 0xda, 0xe2, 0x75, 0x7e, 0x4f, 0x74, 0x01, 0x90, 0x5a,
			0xca, 0xfa, 0xaa, 0xe3, 0xe4, 0xd5, 0x9b, 0x34, 0x9a, 0xdf, 0x6a, 0xce, 0xbd, 0x10, 0x19, 0x0d,
			0xfe, 0x48, 0x90, 0xd1, 0xe6, 0x18, 0x8d, 0x0b, 0x04, 0x6d, 0xf3, 0x44, 0x70, 0x6c, 0x63, 0x1e,
		},
	}

	for i, input := range inputs {
		switch keyLength[i] {
		case KeyLength128:
			Nk = KeyLength128
			Nb = BlockSize128
			Nr = NumOfRounds128
		case KeyLength192:
			Nk = KeyLength192
			Nb = BlockSize192
			Nr = NumOfRounds192
		case KeyLength256:
			Nk = KeyLength256
			Nb = BlockSize256
			Nr = NumOfRounds256
		default:
			t.Error("AES key length must be one of 128, 192, 256 bit")
		}
		expanded := make([]byte, BytesOfWords*Nb*(Nr+1))
		keyExpansion(input, expanded)
		if !bytes.Equal(expanded, expected[i]) {
			t.Errorf("[TestKeyExpansion] case %d failed: expanded != expected : '%v' != '%v'", i, expanded, expected[i])
		}
	}
}

func TestMul(t *testing.T) {
	inputs := [][]byte{
		[]byte{0x57, 0x13},
		[]byte{0x57, 0x02},
		[]byte{0x57, 0x04},
		[]byte{0x57, 0x08},
		[]byte{0x57, 0x10},
		[]byte{0x57, 0x83},
	}
	expected := []byte{
		0xfe,
		0xae,
		0x47,
		0x8e,
		0x07,
		0xc1,
	}

	for i, input := range inputs {
		result := mul(input[0], input[1])
		if result != expected[i] {
			t.Errorf("[TestMul] case %d failed: result != expected : '%v' != '%v'", i, result, expected[i])
		}
	}
}

func TestRotWord(t *testing.T) {
	inputs := [][]byte{
		[]byte{0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0x2a, 0x6c, 0x76, 0x05},
	}
	expected := [][]byte{
		[]byte{0xcf, 0x4f, 0x3c, 0x09},
		[]byte{0x6c, 0x76, 0x05, 0x2a},
	}

	for i, input := range inputs {
		rotWord(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestRotWord] case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestSubWord(t *testing.T) {
	inputs := [][]byte{
		[]byte{0xcf, 0x4f, 0x3c, 0x09},
		[]byte{0x6c, 0x76, 0x05, 0x2a},
	}
	expected := [][]byte{
		[]byte{0x8a, 0x84, 0xeb, 0x01},
		[]byte{0x50, 0x38, 0x6b, 0xe5},
	}

	for i, input := range inputs {
		subWord(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestSubWord] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestSubBytes(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x00, 0x10, 0x20, 0x30, 0x40, 0x50, 0x60, 0x70, 0x80, 0x90, 0xa0, 0xb0, 0xc0, 0xd0, 0xe0, 0xf0},
		[]byte{0x19, 0x3d, 0xe3, 0xbe, 0xa0, 0xf4, 0xe2, 0x2b, 0x9a, 0xc6, 0x8d, 0x2a, 0xe9, 0xf8, 0x48, 0x08},
	}
	expected := [][]byte{
		[]byte{0x63, 0xca, 0xb7, 0x04, 0x09, 0x53, 0xd0, 0x51, 0xcd, 0x60, 0xe0, 0xe7, 0xba, 0x70, 0xe1, 0x8c},
		[]byte{0xd4, 0x27, 0x11, 0xae, 0xe0, 0xbf, 0x98, 0xf1, 0xb8, 0xb4, 0x5d, 0xe5, 0x1e, 0x41, 0x52, 0x30},
	}

	for i, input := range inputs {
		subBytes(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestSubBytes] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestShiftRows(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x63, 0xca, 0xb7, 0x04, 0x09, 0x53, 0xd0, 0x51, 0xcd, 0x60, 0xe0, 0xe7, 0xba, 0x70, 0xe1, 0x8c},
		[]byte{0xd4, 0x27, 0x11, 0xae, 0xe0, 0xbf, 0x98, 0xf1, 0xb8, 0xb4, 0x5d, 0xe5, 0x1e, 0x41, 0x52, 0x30},
	}
	expected := [][]byte{
		[]byte{0x63, 0x53, 0xe0, 0x8c, 0x09, 0x60, 0xe1, 0x04, 0xcd, 0x70, 0xb7, 0x51, 0xba, 0xca, 0xd0, 0xe7},
		[]byte{0xd4, 0xbf, 0x5d, 0x30, 0xe0, 0xb4, 0x52, 0xae, 0xb8, 0x41, 0x11, 0xf1, 0x1e, 0x27, 0x98, 0xe5},
	}

	for i, input := range inputs {
		shiftRows(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestShiftRows] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestMixColumns(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x63, 0x53, 0xe0, 0x8c, 0x09, 0x60, 0xe1, 0x04, 0xcd, 0x70, 0xb7, 0x51, 0xba, 0xca, 0xd0, 0xe7},
		[]byte{0xd4, 0xbf, 0x5d, 0x30, 0xe0, 0xb4, 0x52, 0xae, 0xb8, 0x41, 0x11, 0xf1, 0x1e, 0x27, 0x98, 0xe5},
	}
	expected := [][]byte{
		[]byte{0x5f, 0x72, 0x64, 0x15, 0x57, 0xf5, 0xbc, 0x92, 0xf7, 0xbe, 0x3b, 0x29, 0x1d, 0xb9, 0xf9, 0x1a},
		[]byte{0x04, 0x66, 0x81, 0xe5, 0xe0, 0xcb, 0x19, 0x9a, 0x48, 0xf8, 0xd3, 0x7a, 0x28, 0x06, 0x26, 0x4c},
	}

	for i, input := range inputs {
		mixColumns(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestMixColumns] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestAddRoundKey(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34},
		[]byte{0x04, 0x66, 0x81, 0xe5, 0xe0, 0xcb, 0x19, 0x9a, 0x48, 0xf8, 0xd3, 0x7a, 0x28, 0x06, 0x26, 0x4c},
	}
	keys := [][]byte{
		[]byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c},
		[]byte{0xa0, 0xfa, 0xfe, 0x17, 0x88, 0x54, 0x2c, 0xb1, 0x23, 0xa3, 0x39, 0x39, 0x2a, 0x6c, 0x76, 0x05},
	}
	expected := [][]byte{
		[]byte{0x19, 0x3d, 0xe3, 0xbe, 0xa0, 0xf4, 0xe2, 0x2b, 0x9a, 0xc6, 0x8d, 0x2a, 0xe9, 0xf8, 0x48, 0x08},
		[]byte{0xa4, 0x9c, 0x7f, 0xf2, 0x68, 0x9f, 0x35, 0x2b, 0x6b, 0x5b, 0xea, 0x43, 0x02, 0x6a, 0x50, 0x49},
	}

	for i, input := range inputs {
		addRoundKey(input, keys[i])
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestAddRoundKey] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestInvShiftRows(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x7a, 0xd5, 0xfd, 0xa7, 0x89, 0xef, 0x4e, 0x27, 0x2b, 0xca, 0x10, 0x0b, 0x3d, 0x9f, 0xf5, 0x9f},
		[]byte{0x54, 0xd9, 0x90, 0xa1, 0x6b, 0xa0, 0x9a, 0xb5, 0x96, 0xbb, 0xf4, 0x0e, 0xa1, 0x11, 0x70, 0x2f},
	}
	expected := [][]byte{
		[]byte{0x7a, 0x9f, 0x10, 0x27, 0x89, 0xd5, 0xf5, 0x0b, 0x2b, 0xef, 0xfd, 0x9f, 0x3d, 0xca, 0x4e, 0xa7},
		[]byte{0x54, 0x11, 0xf4, 0xb5, 0x6b, 0xd9, 0x70, 0x0e, 0x96, 0xa0, 0x90, 0x2f, 0xa1, 0xbb, 0x9a, 0xa1},
	}

	for i, input := range inputs {
		invShiftRows(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestInvShiftRows] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestInvSubBytes(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0x7a, 0x9f, 0x10, 0x27, 0x89, 0xd5, 0xf5, 0x0b, 0x2b, 0xef, 0xfd, 0x9f, 0x3d, 0xca, 0x4e, 0xa7},
		[]byte{0x54, 0x11, 0xf4, 0xb5, 0x6b, 0xd9, 0x70, 0x0e, 0x96, 0xa0, 0x90, 0x2f, 0xa1, 0xbb, 0x9a, 0xa1},
	}
	expected := [][]byte{
		[]byte{0xbd, 0x6e, 0x7c, 0x3d, 0xf2, 0xb5, 0x77, 0x9e, 0x0b, 0x61, 0x21, 0x6e, 0x8b, 0x10, 0xb6, 0x89},
		[]byte{0xfd, 0xe3, 0xba, 0xd2, 0x05, 0xe5, 0xd0, 0xd7, 0x35, 0x47, 0x96, 0x4e, 0xf1, 0xfe, 0x37, 0xf1},
	}

	for i, input := range inputs {
		invSubBytes(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestInvSubBytes] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}

func TestInvMixColumns(t *testing.T) {
	Nb = BlockSize128

	inputs := [][]byte{
		[]byte{0xbd, 0x6e, 0x7c, 0x3d, 0xf2, 0xb5, 0x77, 0x9e, 0x0b, 0x61, 0x21, 0x6e, 0x8b, 0x10, 0xb6, 0x89},
		[]byte{0xfd, 0xe3, 0xba, 0xd2, 0x05, 0xe5, 0xd0, 0xd7, 0x35, 0x47, 0x96, 0x4e, 0xf1, 0xfe, 0x37, 0xf1},
	}
	expected := [][]byte{
		[]byte{0x47, 0x73, 0xb9, 0x1f, 0xf7, 0x2f, 0x35, 0x43, 0x61, 0xcb, 0x01, 0x8e, 0xa1, 0xe6, 0xcf, 0x2c},
		[]byte{0x2d, 0x7e, 0x86, 0xa3, 0x39, 0xd9, 0x39, 0x3e, 0xe6, 0x57, 0x0a, 0x11, 0x01, 0x90, 0x4e, 0x16},
	}

	for i, input := range inputs {
		invMixColumns(input)
		if !bytes.Equal(input, expected[i]) {
			t.Errorf("[TestInvMixColumns] Case %d failed: input != expected : '%v' != '%v'", i, input, expected[i])
		}
	}
}
