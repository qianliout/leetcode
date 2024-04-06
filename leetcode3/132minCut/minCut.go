package main

import (
	"fmt"
)

func main() {
	fmt.Println(minCut("aab"))
}

func minCut(s string) int {
	ss := []byte(s)
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = i // 初值
	}
	mem := make(map[string]bool)

	for i := 1; i <= len(ss); i++ {
		if isPalindrome(string(ss[:i]), mem) {
			dp[i] = 0
			continue
		}
		for j := i - 1; j >= 0; j-- {
			if isPalindrome(string(ss[j:i]), mem) && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(s)]
}

func isPalindrome(s string, mem map[string]bool) bool {
	// if ss, ok := mem[s]; ok {
	// 	return ss
	// }
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			mem[s] = false
			return false
		}
		left++
		right--
	}
	// mem[s] = true
	return true
}
