package calc

import "testing"

func TestModPow(t *testing.T) {
	type data struct {
		x, n, mod int
	}

	inputs := []data{
		data{2, 2, 3},
		data{10, 5, 30},
		data{7, 3, 2},
		data{12345, 123, 1234},
	}
	expected := []int{
		1, 10, 1, 837,
	}

	for i, input := range inputs {
		x, n, mod := input.x, input.n, input.mod
		result := modPow(x, n, mod)
		if result != expected[i] {
			t.Errorf("[TestModPow] case %d failed: modPow(%d, %d, %d) = %d , but expected '%d'", i, x, n, mod, result, expected[i])
		}
	}
}
