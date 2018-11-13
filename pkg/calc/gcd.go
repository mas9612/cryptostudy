package calc

import (
	"fmt"
)

type gcdOptions struct {
	verbose bool
}

type GcdOption func(*gcdOptions)

func WithVerbose(verbose bool) GcdOption {
	return func(ops *gcdOptions) {
		ops.verbose = verbose
	}
}

func Gcd(a, b int, options ...GcdOption) int {
	opt := gcdOptions{}
	for _, o := range options {
		o(&opt)
	}

	if b == 0 {
		if opt.verbose == true { // Print intermediate calculation
			fmt.Printf("gcd = %d\n", a)
		}
		return a
	} else {
		reminder := a % b
		if opt.verbose == true { // Print intermediate calculation
			fmt.Printf("%d %% %d = %d\n", a, b, reminder)
		}
		return Gcd(b, reminder, WithVerbose(opt.verbose))
	}
}
