package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := os.Args[1:]
	primes := []int{}

	for _, n := range numbers {
		if i, err := strconv.Atoi(n); err == nil {
			if isPrime(i) {
				primes = append(primes, i)
			}
		}
	}

	fmt.Printf("Prime numbers: %v", primes)
}

func isPrime(n int) bool {
	switch n {
	case 2, 3:
		return true
	default:
		if n%2 == 0 || n%3 == 0 || n == 1 {
			return false
		}
	}

	for i, w := 5, 2; i*i <= n; i += w {
		if n%i == 0 {
			return false
		}

		w = 6 - w
	}
	return true
}
