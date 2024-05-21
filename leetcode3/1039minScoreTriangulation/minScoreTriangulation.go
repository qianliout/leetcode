package main

import (
	"math"
)

func main() {

}

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 为啥是倒序呢
	// dp[i][j]= minE[k=i+1-->j-1](dp[i][k]+dp[k][j]+values[i]*values[j]*values[k])
	// 其中 minE 是 累计取小的意思,[k=i+1-->j-1] 表示中间的 k 从 i+1走到 j-1
	// 从中可以看出 i<k，dp[i]是从 dp[k]转移而来，所以 i 要倒序
	// 同理,j>k,dp[i][j]是从dp[i][k]转移而来，所以 j 要正序

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			ans := math.MaxInt32
			for k := i + 1; k < j; k++ {
				b := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
				ans = min(ans, b)
			}
			dp[i][j] = ans
		}

	}
	return dp[0][n-1]
}

func f(values []int) int {
	n := len(values)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	return dfs(values, 0, n-1, mem)
}

func dfs(values []int, i, j int, mem [][]int) int {
	if i+1 >= j {
		// 少于三个边了
		return 0
	}
	if mem[i][j] > 0 {
		return mem[i][j]
	}

	ans := math.MaxInt32
	for k := i + 1; k < j; k++ {
		b := dfs(values, i, k, mem) + dfs(values, k, j, mem) + values[i]*values[j]*values[k]
		ans = min(ans, b)
	}
	mem[i][j] = ans
	return ans
}
