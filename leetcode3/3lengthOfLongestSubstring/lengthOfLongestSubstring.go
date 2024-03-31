package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("qwe"))
	fmt.Println(lengthOfLongestSubstring("11111"))
	fmt.Println(lengthOfLongestSubstring("1111231"))

}

func lengthOfLongestSubstring(s string) int {
	ss := []byte(s)
	if len(ss) <= 1 {
		return len(s)
	}
	ans := 0
	left, right, window := 0, 0, map[byte]int{}
	for right <= len(ss)-1 {
		ch := ss[right]
		window[ch]++
		for window[ch] > 1 {
			window[ss[left]]--
			left++
		}
		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++
	}
	return ans
}
