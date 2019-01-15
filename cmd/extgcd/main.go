package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mas9612/cryptostudy/pkg/calc"
)

func main() {
	a := flag.Int("a", 0, "a of inverse(a,b), doesn't allowed 0")
	b := flag.Int("b", 0, "b of inverse(a,b), doesn't allowed 0")
	v := flag.Bool("v", false, "Print intermediate calculation")
	flag.Parse()
	if *a == 0 || *b == 0 {
		flag.Usage()
		os.Exit(0)
	}
	options := make([]calc.InverseOption, 0)
	if *v {
		options = append(options, calc.SetVerbose(true))
	}
	fmt.Println(calc.Inverse(*a, *b, options...))
}
