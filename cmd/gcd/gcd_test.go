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
		a, b := input[0], input[1]
		result := gcd(a, b)
		if result != expected[i] {
			t.Errorf("[TestGcd] case %d failed: gcd(%d, %d) = '%d' , expected '%d'", i, a, b, result, expected[i])
		}
	}
}
