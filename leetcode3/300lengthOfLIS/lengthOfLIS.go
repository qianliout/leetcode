package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	if len(nums) <= 1 {
		return len(nums)
	}
	dp[0] = 1
	ans := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		ans = Max(ans, dp[i])
	}
	return ans
}
