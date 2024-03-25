package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for j := range nums {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		fmt.Println(nums)
	}
	return i
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
}
