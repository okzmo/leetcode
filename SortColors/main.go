package main

func sortColors(nums []int) {
	count := [3]int{0, 0, 0}

	for i := range nums {
		count[nums[i]] += 1
	}

	idx := 0
	for j := range count {
		for k := 0; k < count[j]; k++ {
			nums[idx] = j
			idx++
		}
	}
}

func main() {
	nums := []int{2, 0, 1}
	sortColors(nums)
}
