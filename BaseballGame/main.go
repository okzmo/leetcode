package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	arr := []int{}
	for _, v := range operations {
		switch v {
		case "C":
			arr = arr[:len(arr)-1]
		case "D":
			arr = append(arr, arr[len(arr)-1]*2)
		case "+":
			arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
		default:
			n, _ := strconv.Atoi(v)
			arr = append(arr, n)
		}
	}

	ans := 0
	for _, v := range arr {
		ans += v
	}

	return ans
}

func main() {
	ops := []string{"1", "C"}
	fmt.Println(calPoints(ops))
}
