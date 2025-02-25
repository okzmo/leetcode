package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var ans []bool

	if len(nums) < 2 {
		ans = append(ans, false)
		return ans
	}

	for i := range l {
		subarray := append([]int{}, nums[l[i]:r[i]+1]...)
		isValid := true
		sort.Ints(subarray)
		diff := subarray[1] - subarray[0]
		for j := 1; j < len(subarray)-1; j++ {
			if subarray[j+1]-subarray[j] != diff {
				isValid = false
				ans = append(ans, false)
				break
			}
		}
		if isValid {
			ans = append(ans, true)
		}
	}

	return ans
}

func main() {
	nums := []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
	l := []int{0, 1, 6, 4, 8, 7}
	r := []int{4, 4, 9, 7, 9, 10}

	fmt.Println(checkArithmeticSubarrays(nums, l, r))
}
