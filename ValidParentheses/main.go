package main

import (
	"fmt"
)

func isValid(s string) bool {
	matching := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	str := []byte(s)
	stack := []byte{}

	for _, br := range str {
		if _, isOpening := matching[br]; isOpening {
			stack = append(stack, br)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		lastOpening := stack[len(stack)-1]
		closingEquivalent := matching[lastOpening]

		if closingEquivalent == br {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	s := "["
	fmt.Println(isValid(s))
}
