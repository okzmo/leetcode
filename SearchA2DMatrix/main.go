package main

import "fmt"

func search(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1

	for l <= r {
		mid := (l + r) / 2
		val := matrix[mid/n][mid%n]

		if val == target {
			return true
		} else if val < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

func main() {
	nums1 := []int{1, 3, 5, 7}
	nums2 := []int{10, 11, 16, 20}
	nums3 := []int{23, 30, 34, 60}
	nums := [][]int{nums1, nums2, nums3}
	fmt.Println(nums)
	fmt.Println(search(nums, 11))
}
