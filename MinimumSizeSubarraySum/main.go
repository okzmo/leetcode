package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	winSum := nums[0]
	minLen := math.MaxInt32

	for l < len(nums) {
		if winSum >= target {
			minLen = min(minLen, (r-l)+1)
			winSum -= nums[l]
			l++
		} else if winSum < target && r+1 < len(nums) {
			r++
			winSum += nums[r]
		} else {
			fmt.Println(winSum, r, l, minLen)
			break
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

func main() {
	numbers := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, numbers))
}
