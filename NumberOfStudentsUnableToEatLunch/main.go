package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	counters := []int{0, 0}

	for i := range students {
		counters[students[i]] += 1
	}

	for j := range sandwiches {
		if counters[sandwiches[j]] == 0 {
			break
		}
		counters[sandwiches[j]] -= 1
	}

	return counters[0] + counters[1]
}

func main() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}

	fmt.Println(countStudents(students, sandwiches))
}
