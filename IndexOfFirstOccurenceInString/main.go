package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i += len(needle) {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(strStr(s, "leeto"))

}
