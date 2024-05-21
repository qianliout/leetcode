package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
}

func longestPalindromeSubseq2(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初值
	for i := range dp {
		dp[i][i] = 1
	}
	// 通过下面的 dfs 知道，需要从 f(i+1) --> f(i),所以必须先算好 i+1
	// 同理，j 的变化值是 从j-->j+1
	// dp[i][j] = max( (dp[i+1][j-1]+2) && (s[i]==s[j]) || min(dp[i][j-1], dp[i+1][j]) )
	// 可以看出 dp[i] 是从dp[i+1]转化得来
	// dp[i][j]是 d[i][j-1] 转化得来
	// 有个技巧：i 第一行参数，就只 看第一个参数
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = max(dp[i][j], dp[i+1][j-1]+2)
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j], dp[i][j])
			}
		}
	}
	return dp[0][n-1]
}

func longestPalindromeSubseq(s string) int {
	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(s))
	}
	return dfs(s, 0, len(s)-1, mem)
}

func dfs(s string, i, j int, mem [][]int) int {
	if i == j {
		return 1
	}
	if i > j {
		return 0
	}
	if mem[i][j] > 0 {
		return mem[i][j]
	}

	if s[i] == s[j] {
		a := dfs(s, i+1, j-1, mem) + 2
		mem[i][j] = a

		return a
	}
	b := max(dfs(s, i+1, j, mem), dfs(s, i, j-1, mem))
	mem[i][j] = b
	return b
}
