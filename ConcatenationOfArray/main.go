package main

import (
	"fmt"
)

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
