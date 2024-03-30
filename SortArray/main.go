package main

import "fmt"

func partition(nums []int, lo int, hi int) int {
	pivot := nums[hi]
	idx := lo - 1

	for i := lo; i < hi; i++ {
		if nums[i] <= pivot {
			idx++
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	idx++
	nums[hi], nums[idx] = nums[idx], nums[hi]

	return idx
}

func qs(nums []int, lo int, hi int) {
	if lo >= hi {
		return
	}

	pivot := partition(nums, lo, hi)

	qs(nums, lo, pivot-1)
	qs(nums, pivot+1, hi)
}

func sortArray(nums []int) []int {
	qs(nums, 0, len(nums)-1)
	return nums
}

func main() {
	nums := []int{3, 2, 1, 4}
	fmt.Println(sortArray(nums))
}
