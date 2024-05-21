package main

import (
	"fmt"
)

func main() {
	fmt.Println(check([]byte("ADDDFDB"), []byte("ABDF")))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	window := make(map[byte]int)
	ri := 0
	le := 0
	ans := ""
	ss := []byte(s)

	for le <= ri && ri < len(s) {
		ch := ss[ri]
		window[ch]++

		// 确认是否需要收缩
		ri++
		for le <= ri && contain(window, need) {
			// 更新答案
			if ans == "" {
				ans = string(ss[le:ri])
			}
			if ans != "" && len(ans) > ri-le {
				ans = string(ss[le:ri])
			}
			window[ss[le]]--
			le++
		}
	}

	return ans
}

func contain(ss, pp map[byte]int) bool {
	for k, ch := range pp {
		if ss[k] < ch {
			return false
		}
	}
	return true
}

func check(s, t []byte) bool {
	s1, t1 := 0, 0
	for s1 < len(s) && t1 < len(t) {
		if s[s1] == t[t1] {
			s1++
			t1++
		} else {
			s1++
		}
	}
	return t1 == len(t)
}
