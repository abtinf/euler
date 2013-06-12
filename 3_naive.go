package main

import "fmt"

func main() {
	var maxprime uint64 = 0
	var prime uint64 =  1
	var number uint64 = 600851475143
	for prime < number {
		prime = nextPrime(prime)
		if number % prime == 0 {
			fmt.Printf("Factor: %d\n", prime)
			maxprime = prime
			number = number / prime
		}
	}

	fmt.Printf("Large prime factor: %d\n", maxprime)
}

func nextPrime(prime uint64) (uint64) {
	if prime >= 0 && prime <= 2 {
		return prime + 1
	}

	for i := prime + 2; ; i += 2 {
		for j := uint64(3); ; j++ {
			if i % j == 0 {
				if i == j {
					return j
				} else {
					break
				}
			}
		}
	}
	return 0
}

