package main

import "fmt"

func canJump(nums []int) bool {
	jump := 0
	goal := len(nums) - 1
	for i := goal - 1; i > -1; i-- {
		if i+nums[i] >= goal {
			jump += 1
			goal = i
		}
	}

	fmt.Println(jump)

	return goal == 0
}

func canJumpII(nums []int) int {
	jump, furthest, reach := 0, 0, 0
	for i, num := range nums[:len(nums)-1] {
		furthest = max(furthest, i+num)
		if i == reach {
			jump, reach = jump+1, furthest
		}
	}
	return jump
}

func main() {
	nums := []int{2, 3, 1, 1, 1}
	res := canJumpII(nums)
	fmt.Println(res)
}
