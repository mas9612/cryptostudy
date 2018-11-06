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
	fmt.Println(gcd(a, b))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		reminder := a % b
		return gcd(b, reminder)
	}
}
