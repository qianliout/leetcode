package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for i := range s1 {
		need[byte(s1[i])]++
	}
	window := make(map[byte]int)
	le, ri := 0, 0
	for le <= ri && ri < len(s2) {
		ch := byte(s2[ri])
		window[ch]++
		// 确定是否收缩
		for ri-le >= len(s1) {
			window[byte(s2[le])]--
			le++
		}
		ri++

		// 确定是否更新答案
		if check(window, need) {
			return true
		}
	}
	return false
}

func check(window, need map[byte]int) bool {
	if len(window) == 0 {
		return false
	}
	for k, ch := range window {
		if need[k] != ch {
			return false
		}
	}
	for k, ch := range need {
		if window[k] != ch {
			return false
		}
	}

	return true
}
