package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))
}

func minFallingPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初值
	for i := 0; i < n; i++ {
		dp[0][i] = grid[0][i]
	}

	for col := 1; col < m; col++ {
		for row := 0; row < n; row++ {
			// 非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。
			lev := math.MaxInt32
			for k := 0; k < n; k++ {
				if k == row {
					continue
				}
				lev = Min(lev, dp[col-1][k])
			}

			dp[col][row] = lev + grid[col][row]
		}
	}

	ans := math.MaxInt32
	for _, ch := range dp[m-1] {
		if ch == math.MaxInt32 {
			continue
		}
		if ans > ch {
			ans = ch
		}
	}
	return ans
}
