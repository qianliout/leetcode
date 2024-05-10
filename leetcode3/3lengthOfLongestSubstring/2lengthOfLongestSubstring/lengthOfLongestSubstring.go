package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	ans := 0
	if s == "" {
		return ans
	}
	wind := make(map[byte]int)
	left, right := 0, 0
	for left <= right && right < len(s) {
		ch := s[right]
		wind[ch]++
		right++
		for left <= right && wind[ch] > 1 {
			wind[s[left]]--
			left++
		}

		ans = Max(ans, right-left)
	}
	return ans
}
