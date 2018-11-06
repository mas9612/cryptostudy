package main

import "testing"

func TestGcd(t *testing.T) {
	inputs := [][]int{
		[]int{6, 2},
		[]int{100, 1200},
		[]int{32, 64},
	}
	expected := []int{
		2, 100, 32,
	}

	for i, input := range inputs {
		result := gcd(input[0], input[1])
		if result != expected[i] {
			t.Errorf("[input != expected : gcd(%d, %d) != '%d' expected '%d'", input[0], input[1], result, expected[i])
		}
	}
}
