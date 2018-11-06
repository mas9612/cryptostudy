package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fs := flag.NewFlagSet("GCD", flag.ExitOnError)
	a := fs.Int("a", 0, "a of gcd(a,b) , don't allowd 0")
	b := fs.Int("b", 0, "b of gcd(a,b), don't allowd 0")
	v := fs.Bool("v", false, "Print intermediate calculation")
	help := fs.Bool("help", false, "Print help and exit")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Failed to parse command line arguments")
		os.Exit(1)
	}
	if *help {
		fs.Usage()
		os.Exit(0)
	}
	if *a == 0 || *b == 0 {
		fs.Usage()
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

type GcdOption func(*gcdOptions)

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
	} else {
		reminder := a % b
		if opt.verbose == true { // Print intermediate calculation
			fmt.Printf("%d %% %d = %d\n", a, b, reminder)
		}
		return gcd(b, reminder, WithVerbose(opt.verbose))
	}
}
