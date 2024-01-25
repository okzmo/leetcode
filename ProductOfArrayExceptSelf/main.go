package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	multiplier := 1
	for i := 0; i < len(nums); i++ {
		res[i] = multiplier
		multiplier *= nums[i]
	}
	multiplier = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= multiplier
		multiplier *= nums[i]
	}
	return res
}

func main() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums))
}
