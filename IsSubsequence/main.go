package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	p1 := 0
	p2 := 0
	for p2 < len(t) {
		if s[p1] == t[p2] {
			p1++
		}
		if p1 == len(s) {
			return true
		}
		p2++
	}

	return false
}

func main() {
	s := "axc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

}
