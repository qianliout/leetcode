package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(rob2(nums[1:]), rob2(nums[:len(nums)-1]))
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([][2]int, len(nums))
	dp[0][1] = nums[0]
	ans := dp[0][1]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = Max(dp[i-1][0]) + nums[i]
		ans = Max(ans, dp[i][0], dp[i][1])
	}
	return ans
}
