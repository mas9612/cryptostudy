package calc

import "fmt"

// InverseOption represents option passed to Inverse method,
type InverseOption func(*inverseOption)

type inverseOption struct {
	verbose bool
}

// SetVerbose set output verbosity.
// If set this true, intermediate calculation is also printed.
func SetVerbose(verbose bool) InverseOption {
	return func(opt *inverseOption) {
		opt.verbose = verbose
	}
}

// Inverse return value of inverse a mod b
// (inverse / gcd) % b = 1
func Inverse(a, b int, options ...InverseOption) int {
	opt := &inverseOption{
		verbose: false,
	}
	for _, o := range options {
		o(opt)
	}

	_, x, _ := ExtGcd(a, b, SetVerbose(opt.verbose))
	for {
		if opt.verbose {
			fmt.Printf("Inverse: %d\n", x)
		}
		if x >= 0 {
			break
		}
		x += b
	}
	return x
}

// ExtGcd return value:
// - gcd, x, y
func ExtGcd(a, b int, options ...InverseOption) (int, int, int) {
	opt := &inverseOption{
		verbose: false,
	}
	for _, o := range options {
		o(opt)
	}

	gcd0, gcd1 := a, b
	a0, a1 := 1, 0
	b0, b1 := 0, 1
	if opt.verbose {
		fmt.Printf("%d * %d + %d * %d = %d\n", a0, gcd0, b0, gcd1, gcd0)
		fmt.Printf("%d * %d + %d * %d = %d\n", a1, gcd0, b1, gcd1, gcd1)
	}

	for gcd1 != 0 {
		reminder := gcd0 % gcd1
		quotient := gcd0 / gcd1
		gcd0, gcd1 = gcd1, reminder
		if opt.verbose && reminder != 0 {
			fmt.Printf("(%d - %d * %d) * %d + (%d - %d * %d) * %d = ", a0, quotient, a1, a, b0, quotient, b1, b)
		}
		a0, a1 = a1, a0-quotient*a1
		b0, b1 = b1, b0-quotient*b1
		if opt.verbose && reminder != 0 {
			fmt.Printf("%d * %d + %d * %d = %d\n", a1, a, b1, b, reminder)
		}
	}
	fmt.Print("\n")
	return gcd0, a0, b0
}
