package main

import "fmt"

func canJump(nums []int) bool {
	goal := len(nums) - 1
	for i := goal - 1; i > -1; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}

func main() {
	nums := []int{0, 1}
	res := canJump(nums)
	fmt.Println(res)
}
