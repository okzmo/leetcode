package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxVal := 0

	for l < r {
		size := r - l
		temp := min(height[l], height[r]) * size
		if temp > maxVal {
			maxVal = temp
		}

		if height[l] < height[r] {
			l++
		} else if height[l] > height[r] {
			r--
		} else {
			l++
			r--
		}
	}

	return maxVal
}

func main() {
	numbers := []int{7, 3, 8, 4, 5, 2, 6, 8, 1}
	fmt.Println(maxArea(numbers))

}
