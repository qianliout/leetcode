package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := range dp {
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1], dp[i][j])
			}
		}
	}
	return dp[0][len(s)-1]
}
