package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(minDistance("horse", "hors"))
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				/*
					插入一个字符 abc由ab转化得来 ,插入c，对应：dp[i][j] = dp[i][j-1]+1
					删除一个字符 ab由abc转化得来，删除 c ，对应 dp[i][j] = dp[i-1][j]+1
					替换一个字符 abc由abd转化得来，把d变成c，对应 dp[i][j] = dp[i-1][j-1]+1
				*/
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
