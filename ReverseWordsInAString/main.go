package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	spl := strings.Fields(strings.TrimSpace(s))
	for i, j := 0, len(spl)-1; i < j; i, j = i+1, j-1 {
		spl[i], spl[j] = spl[j], spl[i]
	}
	return strings.Join(spl, " ")
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}
