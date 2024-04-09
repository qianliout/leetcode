package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
}

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}

	k = k % len(prices)
	if k == 0 {
		k = len(prices)
	}
	dp := make([][][2]int, len(prices))
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}
	// 初值
	for j := 0; j <= k; j++ {
		dp[0][j][0] = 0
		dp[0][j][1] = -prices[0]
	}
	ans := 0
	for i := 1; i < len(prices); i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][1] = Max(dp[i-1][j-1][0]-prices[i], dp[i-1][j][1])
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			ans = Max(ans, dp[i][j][0])
		}
	}
	return ans
}
