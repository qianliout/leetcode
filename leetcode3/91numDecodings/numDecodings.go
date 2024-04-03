package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("12"))
}

func numDecodings(s string) int {
	nums := make([]int, 0)
	for i := range s {
		nums = append(nums, int(s[i])-48)
	}
	if len(nums) == 0 || nums[0] == 0 {
		return 0
	}

	// dp[i] 表示 nums[:i)的解法
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[i-1] == 1 || nums[i-1] == 2 {
				dp[i+1] = dp[i-1]
				continue
			}
			return 0
		}

		if nums[i-1] == 1 || (nums[i-1] == 2 && nums[i] <= 6 && nums[i] >= 0) {
			dp[i+1] = dp[i-1] + dp[i]
			continue
		}

		dp[i+1] = dp[i]
	}
	return dp[len(nums)]
}
