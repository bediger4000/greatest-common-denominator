package main

/*
Prime number generator:
https://www.youtube.com/watch?v=hB05UFqOtFA&feature=youtu.be

"Advanced Topics in Programming Languages: Concurrency/message passing
Newsqueak", 2012 Rob Pike talk.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []int64
	var maxNumber int64
	for _, str := range os.Args[1:] {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, n)
		if n > maxNumber {
			maxNumber = n
		}
	}

	// create map, keys prime factors, values are
	// how many of the n numbers the prime is a factor of
	factors := make(map[int64]int)
	prime := make(chan int64)
	go sieve(prime)
	for i := 0; true; i++ {
		p := <-prime
		if p > maxNumber {
			break
		}

		for _, n := range numbers {
			if n%p == 0 {
				factors[p]++
			}
		}
	}

	var gcd int64 = 1
	for prime, count := range factors {
		if count == len(numbers) {
			// prime is a factor of all n numbers
			// find minimum count of times prime is a factor
			// of all n numbers
			var instances []int64
			for _, n := range numbers {
				var c int64 = 0
				for n%prime == 0 {
					n /= prime
					c++
				}
				instances = append(instances, c)
			}

			min := instances[0]
			for _, i := range instances[1:] {
				if i < min {
					min = i
				}
			}

			for ; min > 0; min-- {
				gcd *= prime
			}
		}
	}
	fmt.Printf("%d\n", gcd)
}

func counter(c chan int64) {
	var i int64 = 1
	for {
		i++
		c <- i
	}
}

func filter(prime int64, recv chan int64, send chan int64) {
	for {
		i := <-recv
		if i%prime != 0 {
			// fmt.Printf("filter for %d passing along %d\n", prime, i)
			send <- i
		}
	}
}

func sieve(prime chan int64) {
	c := make(chan int64)
	go counter(c)
	for {
		p := <-c
		prime <- p
		newc := make(chan int64)
		go filter(p, c, newc)
		c = newc
	}
}
