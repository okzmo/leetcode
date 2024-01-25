package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	start, total, tank := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		total += gas[i] - cost[i]

		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}

	if total < 0 {
		return -1
	}
	return start
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
