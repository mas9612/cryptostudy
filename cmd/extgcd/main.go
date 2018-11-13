package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mas9612/cryptostudy/pkg/calc"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: inv NUMA NUMB")
		os.Exit(1)
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(calc.Inverse(a, b))
}
