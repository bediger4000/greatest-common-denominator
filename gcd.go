package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	current, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, str := range os.Args[2:] {
		m, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		current = gcd(current, m)
	}

	fmt.Printf("%d\n", current)
}

// Euclid's algorithm for finding Greatest Common Divisor
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
