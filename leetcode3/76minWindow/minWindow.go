package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	ss, tt := []byte(s), []byte(t)
	window := make(map[byte]int)
	tm := make(map[byte]int)
	for i := range tt {
		tm[tt[i]]++
	}
	ans := ""
	i, left := 0, 0

	for i < len(ss) {
		window[ss[i]]++
		for full(window, tm) {
			if ans == "" || len(ans) >= i-left+1 {
				ans = string(ss[left : i+1])
			}
			window[ss[left]]--
			left++
		}
		i++
	}
	return ans
}

func full(used map[byte]int, tt map[byte]int) bool {
	for i, v := range tt {
		if used[i] < v {
			return false
		}
	}
	return true
}
