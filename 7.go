package main

import "fmt"
import "errors"

func main() {
	index := 10001;
	prime := 0
	primes, _ := Sieve(10000000)

	for _, v := range primes {
		if v == 0 {
			continue
		}
		if index == 0 {
			prime = v
			break
		}
		index--
	}

	fmt.Printf("Prime %d: %d\n", index, prime)
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
