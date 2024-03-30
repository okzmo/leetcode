package main

import "fmt"

func isBadVersion(version int) bool {
	return version > 3
}

func firstBadVersion(n int) int {
	lo, hi := 1, n

	for lo <= hi {
		mid := (lo + hi) / 2
		if isBadVersion(mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func main() {
	fmt.Println(firstBadVersion(5))
}
