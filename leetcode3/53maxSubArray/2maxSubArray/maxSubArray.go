package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = Max(nums[i], dp[i-1]+nums[i])
		ans = Max(ans, dp[i])
	}

	return ans
}
