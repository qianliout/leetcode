package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}
	m, n := len(dungeon)-1, len(dungeon[0])-1
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 初值
	// dp[i][j]表示走到 i,j 这个位置刚好死去至少需要补充的血量
	// dungeon 认为是消耗的血量,那么 Max(0, -dungeon[m][n]) 就是要补充的血量

	dp[m][n] = Max(0, -dungeon[m][n])

	// 最左边的一排
	for i := m - 1; i >= 0; i-- {
		dp[i][n] = Max(dp[i+1][n]-dungeon[i][n], 0)
		// 对于 Max(dp[i+1][n]-dungeon[i][n], 0)的理解
		// dp[i+1][n] 是 [i+1,n] 这个点要补充的
		// -dungeon[i][n] 是i,n 这个点要补充的，两个之和就是从[i,n]走到[i+1][n]这个点，在开始处要补充的血量
	}

	for j := n - 1; j >= 0; j-- {
		dp[m][j] = Max(dp[m][j+1]-dungeon[m][j], 0)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = Max(Min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 0)
		}
	}
	return dp[0][0] + 1
}
