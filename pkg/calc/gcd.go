package calc

import (
	"fmt"
)

type gcdOptions struct {
	verbose bool
}

// GcdOption is a function to realize variable arguments
type GcdOption func(*gcdOptions)

// WithVerbose is ops.verbose
func WithVerbose(verbose bool) GcdOption {
	return func(ops *gcdOptions) {
		ops.verbose = verbose
	}
}

// Gcd return gcd(a,b)
// you can set WithVerbose when you need display details
func Gcd(a, b int, options ...GcdOption) int {
	opt := gcdOptions{}
	for _, o := range options {
		o(&opt)
	}

	if b == 0 {
		// Print intermediate calculation
		if opt.verbose == true {
			fmt.Printf("gcd = %d\n", a)
		}
		return a
	}
	reminder := a % b
	// Print intermediate calculation
	if opt.verbose == true {
		fmt.Printf("%d %% %d = %d\n", a, b, reminder)
	}
	return Gcd(b, reminder, WithVerbose(opt.verbose))
}
