package main

import (
	"bytes"
	"fmt"
)

func main() {
	// fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
	// fmt.Println(findTheLongestSubstring("leetcodeisgreat"))
	fmt.Println(findTheLongestSubstring("bcbcbc"))
}

// 能得到结果，但是会  超时
func findTheLongestSubstring1(s string) int {
	y := []byte("aeiou")

	preSum := make([][5]int, len(s)+1)
	for i := range s {
		preSum[i+1] = preSum[i]
		idx := bytes.LastIndexByte(y, s[i])
		if idx == -1 {
			continue
		}
		preSum[i+1][idx]++
	}

	ans := 0
	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i; j++ {
			find := false
			for k := 0; k < 5; k++ {
				if (preSum[i][k]-preSum[j][k])%2 != 0 {
					find = true
					break
				}
			}
			if !find {
				ans = max(ans, i-j)
			}
		}
	}
	return ans
}

func findTheLongestSubstring(s string) int {
	pos := make([]int, 1<<5+1)
	for i := range pos {
		pos[i] = -1
	}
	status := 0
	ans := 0
	y := []byte("aeiou")
	pos[0] = 0

	for i := range s {
		idx := bytes.LastIndexByte(y, s[i])
		/*
						这样写会出问题，因为元音字母出现0次也是答案
			if idx==-1 {
				continue
			}
		*/

		if idx != -1 {
			status = status ^ (1 << idx)
		}

		if pos[status] >= 0 {
			ans = max(ans, i+1-pos[status])
		} else {
			pos[status] = i + 1
		}
	}
	return ans
}
