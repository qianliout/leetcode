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

	return Max(rob2(nums[:len(nums)-1]), rob2(nums[1:]))
}

func rob3(nums []int) int {
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

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp1 := make([]int, len(nums)+1) // rob
	dp2 := make([]int, len(nums)+1) // not rob

	dp1[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp1[i] = Max(dp2[i-1] + nums[i])
		dp2[i] = Max(dp1[i-1], dp2[i-1])
	}

	return Max(dp2[len(nums)-1], dp1[len(nums)-1])
}
