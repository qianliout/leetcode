package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	ans := 0
	// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
	for i := 1; i < len(prices); i++ {
		if i-1 == 0 {
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
		}

		if i-1 > 0 {
			dp[i][1] = Max(dp[i-1][1], dp[i-2][0]-prices[i])
			dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		}
		ans = Max(ans, dp[i][0])

	}
	return ans
}
