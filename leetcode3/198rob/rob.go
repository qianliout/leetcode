package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([][2]int, len(nums))
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = Max(dp[i-1][0] + nums[i])
	}
	return Max(dp[n-1][0], dp[n-1][1])
}
