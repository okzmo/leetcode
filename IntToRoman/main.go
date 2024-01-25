package main

import (
	"fmt"
	"strings"
)

type RomInt struct {
	rom string
	val int
}

var Numbers = []RomInt{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func intToRoman(num int) string {
	var res []string

	for _, v := range Numbers {
		cNum := num / v.val
		fmt.Println(cNum)
		res = append(res, strings.Repeat(v.rom, cNum))
		num = num % v.val
		fmt.Println(num)
	}

	return strings.Join(res, "")
}

func main() {
	num := 1994
	fmt.Println(intToRoman(num))
}
