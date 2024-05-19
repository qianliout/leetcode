package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	mem := make([][]int, m)
	for i := range mem {
		mem[i] = make([]int, n)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	return dfs([]byte(text1), []byte(text2), m-1, n-1, mem)
}

// 从后向前
func f(text1 []byte, text2 []byte) int {

	t1, t2 := len(text1), len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}
	// 技巧，多开一个数组，减少数组越界的判断
	for k := t1; k > 0; k-- {
		for j := t2; j > 0; j-- {
			if text1[k-1] == text2[j-1] {
				dp[k][j] = max(dp[k][j], dp[k-1][j-1]+1)
			} else {
				dp[k][j] = max(dp[k][j-1], dp[k-1][j], dp[k-1][j-1])
			}
		}
	}
	return dp[t1][t2]
}

// 从前向后
func f2(text1, text2 string) int {

	t1, t2 := len(text1), len(text2)
	dp := make([][]int, t1+1)
	for i := range dp {
		dp[i] = make([]int, t2+1)
	}

	for i := 1; i <= t1; i++ {
		for j := 1; j <= t2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[t1][t2]
}

// 必须加上记忆化
func dfs(text1 []byte, text2 []byte, i, j int, mem [][]int) int {
	if i < 0 || j < 0 {
		return 0
	}
	if mem[i][j] >= 0 {
		return mem[i][j]
	}

	if text1[i] == text2[j] {
		a := dfs(text1, text2, i-1, j-1, mem) + 1
		mem[i][j] = a
		return a
	}
	a := max(dfs(text1, text2, i, j-1, mem), dfs(text1, text2, i-1, j, mem))
	mem[i][j] = a
	return a
}
