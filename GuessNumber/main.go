package main

import "fmt"

func guess(num int) int {
	if num < 6 {
		return 1
	} else if num > 6 {
		return -1
	} else {
		return 0
	}
}

func guessNumber(n int) int {
	lo, hi := 0, n

	for lo <= hi {
		mid := (lo + hi) / 2
		if guess(mid) < 0 {
			hi = mid - 1
		} else if guess(mid) > 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	fmt.Println(guessNumber(10))
}
