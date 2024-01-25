package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")

	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	sentence := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(sentence))
}
