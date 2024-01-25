package main

import "fmt"

func romanToInt(s string) int {
	res := 0
	romanNumber := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanNumber[s[i]] < romanNumber[s[i+1]] {
			res += romanNumber[s[i+1]] - romanNumber[s[i]]
			i++
		} else {
			res += romanNumber[s[i]]
		}
	}

	return res
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
