package main

import (
	"fmt"
	"strings"
)

func mapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, val := range a {
		if b[key] != val {
			return false
		}
	}

	return true
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) == 0 {
		return []int{}
	}

	var result []int
	windowSize := len(strings.Join(words, ""))
	wordCount := make(map[string]int)
	wordLen := len(words[0])

	for _, word := range words {
		wordCount[word]++
	}

	for i := 0; i <= len(s)-windowSize; i++ {
		window := s[i : windowSize+i]

		chunkCount := make(map[string]int)
		for j := 0; j < windowSize; j += wordLen {
			chunk := window[j : j+wordLen]
			chunkCount[chunk]++
		}

		if mapsEqual(chunkCount, wordCount) {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	input := "ababaab"
	words := []string{"ab", "ba", "ba"}
	// input2 := "wordgoodgoodgoodbestword"
	// words2 := []string{"word", "good", "best", "good"}
	// input3 := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	// words3 := []string{"fooo", "barr", "wing", "ding", "wing"}
	//
	fmt.Println(findSubstring(input, words))
	// fmt.Println(findSubstring(input2, words2))
	// fmt.Println(findSubstring(input3, words3))
}
