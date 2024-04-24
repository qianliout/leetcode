package main

import (
	"fmt"
)

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams2(s string, p string) []int {
	window := make(map[int32]int)
	need := make(map[int32]int)
	for _, ch := range p {
		need[ch]++
	}
	ans := make([]int, 0)
	le, ri := 0, 0

	for le <= ri && ri < len(s) {
		ch := int32(s[ri])
		ri++

		// 窗口数据更新
		window[ch]++

		// 更新答案
		if equal(window, need) {
			ans = append(ans, le)
		}

		// 判断是否需要收缩
		// 对于滑动窗口的题，因为每次移动一个，所以基本上都是直接使用一个 if 搞定
		if ri >= len(p) {
			window[int32(s[le])]--
			le++
		}
	}

	return ans
}

func findAnagrams(s string, p string) []int {
	window := make(map[int32]int)
	need := make(map[int32]int)
	for _, ch := range p {
		need[ch]++
	}
	ans := make([]int, 0)
	le, ri := 0, 0

	for le <= ri && ri < len(s) {
		ch := int32(s[ri])
		// 窗口数据更新
		window[ch]++
		// 判断是否收缩窗口
		// 这里要注意,在判断前，ri已经增加了
		for ri-le >= len(p) {
			window[int32(s[le])]--
			le++
		}
		ri++
		// 更新答案
		if equal(window, need) {
			ans = append(ans, le)
		}
	}
	return ans
}

func equal(ss, pp map[int32]int) bool {
	if len(ss) == 0 {
		return false
	}
	for k, ch := range ss {
		if pp[k] != ch {
			return false
		}
	}
	for k, ch := range pp {
		if ss[k] != ch {
			return false
		}
	}

	return true
}
