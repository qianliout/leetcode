package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func rob(nums []int) int {
	dp := make([][2]int, len(nums))
	dp[1][1] = nums[0]
	ans := dp[1][1]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = Max(dp[i-1][0] + nums[i])
		ans = Max(ans, dp[i][1])
	}
	return ans
}
