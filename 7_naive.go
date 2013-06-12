package main

import "fmt"

func main() {
	index := 10001;
	prime := 1

	for i := 1; i <= 10001; i++ {
		prime = nextPrime(prime)
	}

	fmt.Printf("Prime %d: %d\n", index, prime)
}

func nextPrime(prime int) (int) {
	if prime >= 0 && prime <= 2 {
		return prime + 1
	}

	for i := prime + 2; ; i += 2 {
		for j := 3; ; j++ {
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

