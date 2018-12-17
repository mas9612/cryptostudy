package calc

import (
	"testing"
)

func TestExtGcd(t *testing.T) {
	inputs := [][]int{
		[]int{10, 340},
		[]int{3, 333},
		[]int{6, 8},
		[]int{12, 44},
		[]int{1234, 1233},
	}
	gcds := []int{
		10, 3, 2, 4, 1,
	}
	for i, input := range inputs {
		a, b := input[0], input[1]
		gcd, x, y := ExtGcd(a, b)
		if gcd != gcds[i] {
			t.Errorf("[TestExtGcd] case %d failed: gcd(%d, %d) != %d , expected : %d", i, a, b, gcd, gcds[i])
		}
		if (a*x+b*y)/gcd != 1 {
			t.Errorf("[TestExtGcd] case %d failed: (%dx + %dy)/gcd(%d,%d) != 1 , result: %d", i, a, b, a, b, (a*x-b*y)/gcd)
		}
	}
}

func TestInverse(t *testing.T) {
	inputs := [][]int{
		[]int{10, 340},
		[]int{3, 333},
		[]int{6, 8},
		[]int{12, 44},
		[]int{1234, 1233},
	}
	gcds := []int{
		10, 3, 2, 4, 1,
	}
	for i, input := range inputs {
		a, b := input[0], input[1]
		result := Inverse(a, b)
		if (a*result)%b != gcds[i] {
			t.Errorf("[TestInverse] case %d failed: %d*inverse(%d, %d) mod %d != %d ", i, a, a, b, b, gcds[i])
		}
	}
}
