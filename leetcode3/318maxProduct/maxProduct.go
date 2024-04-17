package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func maxProduct2(words []string) int {
	wm := make([]map[byte]int, len(words))
	for i := range words {
		wm[i] = cov(words[i])
	}

	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if notSame(wm[i], wm[j]) {
				ans = Max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func maxProduct(words []string) int {
	wm := make([]int, len(words))
	for i := range words {
		wm[i] = cov2(words[i])
	}

	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if wm[i]&wm[j] == 0 {
				ans = Max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func cov2(s string) int {
	ans := 0
	for i := range s {
		ans = ans | (1 << (s[i] - 'a'))
	}
	return ans
}

func cov(s string) map[byte]int {
	ans := make(map[byte]int)
	for i := range s {
		ans[s[i]]++
	}
	return ans
}

func notSame(be, af map[byte]int) bool {
	for k := range be {
		if af[k] > 0 {
			return false
		}
	}
	return true
}
