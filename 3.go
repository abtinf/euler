package main

import "fmt"
import "errors"

func main() {
	var maxprime uint64 = 0
	var number uint64 = 600851475143
	primes, _ := Sieve(10000)

	for _, v := range primes {
		if v == 0 {
			continue
		}
		if number % uint64(v) == 0 {
			fmt.Printf("Factor: %d\n", v)
			number = number / uint64(v)
			maxprime = uint64(v)
		}
	}

	fmt.Printf("Large prime factor: %d\n", maxprime)
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
