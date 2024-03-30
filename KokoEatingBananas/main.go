package main

import (
	"fmt"
	"math"
	"slices"
)

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 0, slices.Max(piles)

	for lo <= hi {
		mid := (lo + hi) / 2
		hours := 0

		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(mid)))
		}

		if hours <= h {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func main() {
	nums := []int{3, 6, 7, 11}
	fmt.Println(minEatingSpeed(nums, 8))
}
