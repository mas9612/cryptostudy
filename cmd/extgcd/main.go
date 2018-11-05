package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	fmt.Println(inverse(a, b))
}

func inverse(a, b int) int {
	_, x, _ := extGcd(a, b)
	for x < 0 {
		x += b
	}
	return x
}

// return value:
// - gcd, x, y
func extGcd(a, b int) (int, int, int) {
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
