package main

func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	majorityElement(nums)
}
