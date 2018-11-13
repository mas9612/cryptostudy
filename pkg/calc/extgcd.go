package calc

// Inverse return value of inverse a mod b
// (inverse / gcd) % b = 1
func Inverse(a, b int) int {
	_, x, _ := ExtGcd(a, b)
	for x < 0 {
		x += b
	}
	return x
}

// ExtGcd return value:
// - gcd, x, y
func ExtGcd(a, b int) (int, int, int) {
	gcd0, gcd1 := a, b
	a0, a1 := 1, 0
	b0, b1 := 0, 1

	for gcd1 != 0 {
		reminder := gcd0 % gcd1
		quotient := gcd0 / gcd1

		gcd0, gcd1 = gcd1, reminder
		a0, a1 = a1, a0-quotient*a1
		b0, b1 = b1, b0-quotient*b1
	}
	return gcd0, a0, b0
}
