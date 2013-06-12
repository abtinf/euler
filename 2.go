package main

import "fmt"

func main() {
	prev := 1
	cur := 2
	total := 0

	for cur < 4000000 {
		if cur % 2 == 0 {
			total += cur
		}
		next := prev + cur
		prev = cur
		cur = next
	}

	fmt.Printf("Sum: %d\n", total)
}
