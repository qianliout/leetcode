package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(25))
	fmt.Println(numSquares(6665))
}

func numSquares(n int) int {
	dp := make(map[int]int)
	dp[0], dp[1] = 0, 1

	for i := 1; i <= n; i++ {
		dp[i] = i // 初值
		for j := 1; j*j <= i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}
