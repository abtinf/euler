package main

import "fmt"

func main() {
	max := 20
	nums := make([]int, max)

	for i := 1; i <= max; i++ {
		nums[i-1] = i
	}

	lcm := BruteLcm(nums)

	fmt.Printf("LCM of 1-%d is: %d\n", max, lcm)

}

func BruteLcm(nums []int) (result int) {
	result = 0
	for {
		result++
		match := true
		for _, v := range nums {
			if result % v != 0 {
				match = false
				break
			}
		}
		if match {
			break
		}
	}
	return result
}
