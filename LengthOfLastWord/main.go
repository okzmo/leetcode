package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	spl := strings.Fields(strings.TrimSpace(s))
	return len(spl[len(spl)-1])
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
