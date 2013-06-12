package main

import "fmt"

func main() {
	max := 100
	nums := make([]int, max)
	for i := 1; i <= max; i++ {
		nums[i-1] = i
	}
	sumOfSquares := SumOfSquares(nums)
	squareOfSums := SquareOfSums(nums)
	difference := squareOfSums - sumOfSquares
	fmt.Printf("Square of Sums - Sum of Squares for 1-100: %d\n", difference)
}

func SumOfSquares(nums []int) (sum int) {
	for _, v := range nums {
		sum += v * v
	}
	return sum
}

func SquareOfSums(nums []int) (sum int) {
	for _, v := range nums {
		sum += v
	}
	return sum * sum
}
