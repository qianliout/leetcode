package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum(matrix))

}

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	// dp[i][j]  表示到matrix[i][j]最小的结果合
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初值
	dp[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}
	for col := 1; col < m; col++ {
		for row := 0; row < n; row++ {
			n1 := dp[col-1][row]
			if row-1 >= 0 {
				n1 = Min(n1, dp[col-1][row-1])
			}
			if row+1 < n {
				n1 = Min(n1, dp[col-1][row+1])
			}

			dp[col][row] = n1 + matrix[col][row]
		}
	}
	ans := dp[m-1][0]
	for _, ch := range dp[m-1] {
		if ans > ch {
			ans = ch
		}
	}
	return ans
}
