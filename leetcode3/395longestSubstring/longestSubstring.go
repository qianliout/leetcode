package main

import (
	"bytes"
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(longestSubstring("ababbc", 2))
}

func longestSubstring(s string, k int) int {
	return dfs([]byte(s), k)
}

func dfs(dd []byte, k int) int {
	if len(dd) < k {
		return 0
	}
	mem := make(map[byte]int)
	for _, by := range dd {
		mem[by]++
	}

	for ch, fre := range mem {
		if fre < k {
			ma := 0
			split := bytes.Split(dd, []byte{ch})
			for _, sp := range split {
				ma = Max(ma, dfs(sp, k))
			}
			return ma
		}
	}
	return len(dd)
}
