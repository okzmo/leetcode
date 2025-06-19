package main

func lengthOfLongestSubstring(s string) int {
	l, r, ans := 0, 0, 0
	hash := make(map[byte]int)

	for r = range s {
		if _, ok := hash[s[r]]; ok {
			l = max(l, hash[s[r]]+1)
		}

		hash[s[r]] = r
		ans = max(ans, r-l+1)
	}

	return ans
}

func main() {
	lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("bbbbb")
	lengthOfLongestSubstring("pwwkew")
	lengthOfLongestSubstring("dvdf")
	lengthOfLongestSubstring(" ")
}
