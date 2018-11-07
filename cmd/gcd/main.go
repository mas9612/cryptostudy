package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	a := flag.Int("a", 0, "a of gcd(a,b) , don't allowd 0")
	b := flag.Int("b", 0, "b of gcd(a,b), don't allowd 0")
	v := flag.Bool("v", false, "Print intermediate calculation")
	help := flag.Bool("help", false, "Print help and exit")
	flag.Parse()
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if *a == 0 || *b == 0 {
		flag.Usage()
		os.Exit(1)
	}
	if *v == true {
		fmt.Println(gcd(*a, *b, WithVerbose(true)))
	} else {
		fmt.Println(gcd(*a, *b, WithVerbose(false)))
	}
}

type gcdOptions struct {
	verbose bool
}

// GcdOption is a method to set appropriate option to gcdOptions struct
type GcdOption func(*gcdOptions)

// WithVerbose set option to increase verbosity in gcd()
func WithVerbose(verbose bool) GcdOption {
	return func(ops *gcdOptions) {
		ops.verbose = verbose
	}
}

func gcd(a, b int, options ...GcdOption) int {
	opt := gcdOptions{}
	for _, o := range options {
		o(&opt)
	}

	if b == 0 {
		if opt.verbose == true { // Print intermediate calculation
			fmt.Printf("gcd = %d\n", a)
		}
		return a
	}
	reminder := a % b
	if opt.verbose == true { // Print intermediate calculation
		fmt.Printf("%d %% %d = %d\n", a, b, reminder)
	}
	return gcd(b, reminder, WithVerbose(opt.verbose))
}
