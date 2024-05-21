package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	matri := [][]byte{[]byte("10100"), []byte("10111"), []byte("11111"), []byte("10010")}
	fmt.Println(maximalSquare(matri))
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			ans = 1
			dp[i][0] = 1
		}
	}
	for j := 0; j < n; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ans = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			} else {
				dp[i][j] = Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return ans * ans
}
