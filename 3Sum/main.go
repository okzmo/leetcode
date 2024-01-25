package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	var sum, l, r int
	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}

		l, r = i+1, len(nums)-1
		for l < r {
			sum = v + nums[l] + nums[r]
			if sum < 0 {
				l++
			} else if sum > 0 {
				r--
			} else {
				res = append(res, []int{v, nums[l], nums[r]})
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return res
}

func main() {
	numbers := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(numbers))

}
