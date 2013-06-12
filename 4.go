package main

import "fmt"

func main() {
	var max int

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			if i * j > max && IsPalindrome(i*j) {
				max = i * j
			}
		}
	}

	fmt.Printf("Largest palindrome: %d\n", max)
}

func IsPalindrome(n int) (bool) {
	original := n
	reverse := 0

	for n > 0 {
		reverse = reverse * 10 + n % 10
		n = n / 10
	}

	return original == reverse
}
