package main

import "fmt"

// secondLast and last == steps
// secondLast will be the floor and last the first step (at the beginning)
//
// We do that because if we try to check how many ways we can climb 3 stairs there's 3 ways
// and for two stairs there's 2 and for 1 stair there's 1. So if we take the secondLast and last we get = 3.
//
// If we try 4 stairs which have 5 ways. We check for 3 stairs there's 3 ways and for two stairs 2 ways, 3 + 2 = 5.

func climbStairs(n int) int {
	secondLast, last := 0, 1
	for i := 0; i < n; i++ {
		secondLast, last = last, secondLast+last
	}

	return last
}

func main() {
	fmt.Println(climbStairs(4))
}
