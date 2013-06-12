package main

import "fmt"
import "errors"

func main() {
	max := 2000000

	primes, _ := Sieve(max)

	var sum uint64 = 0

	for _, v := range primes {
		sum += uint64(v)
	}

	fmt.Printf("Sum of primes less than 2 million: %d\n", sum)
}

func Sieve(max int) ([]int, error) {
	if max < 2 {
		return nil , errors.New("max must be >= 2")
	}

	num := make([]int, max)
	for i := 2; i < max; i++ {
		num[i] = i
	}

	for i := 2; i < max; {
		for j := i; ; {
			j += i
			if j >= max {
				break
			}
			num[j] = 0
		}
		i++
		for i < max && num[i] == 0 {
			i++
		}

	}

	return num, nil
}
