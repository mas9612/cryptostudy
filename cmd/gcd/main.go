package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mas9612/cryptostudy/pkg/calc"
)

func main() {
	a := flag.Int("a", 0, "a of gcd(a,b) , don't allowed 0")
	b := flag.Int("b", 0, "b of gcd(a,b), don't allowed 0")
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
		fmt.Println(calc.Gcd(*a, *b, calc.WithVerbose(true)))
	} else {
		fmt.Println(calc.Gcd(*a, *b, calc.WithVerbose(false)))
	}
}
