package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	hIndex := 0
	for i, v := range citations {
		if i+1 <= v {
			hIndex = i + 1
		} else {
			return hIndex
		}
	}

	return hIndex
}

func main() {
	citations := []int{1, 3, 4, 5, 6}
	fmt.Println(hIndex(citations))

}
