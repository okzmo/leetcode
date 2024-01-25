package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var res []byte
	increment := 2 * (numRows - 1)
	for r := 0; r < numRows; r++ {
		for i := r; i < len(s); i += increment {
			res = append(res, s[i])
			j := i + increment - 2*r
			if r > 0 && r < numRows-1 && j < len(s) {
				res = append(res, s[j])
			}
		}
	}
	return string(res)
}

func main() {
	s := "PAYPALISHIRING"
	fmt.Println(convert(s, 4))
}
