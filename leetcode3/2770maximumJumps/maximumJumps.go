package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 2))
	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 0))
	fmt.Println(maximumJumps([]int{0, 2, 1, 3}, 1))
}

func maximumJumps(nums []int, target int) int {
	if len(nums) <= 1 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n+1)
	for j := 1; j < n; j++ {
		// 初值是关键
		dp[j] = math.MinInt
		for i := j - 1; i >= 0; i-- {
			/*
			   0 <= i < j < n
			   -target <= nums[j] - nums[i] <= target
			*/
			if nums[j]-nums[i] <= target && nums[j]-nums[i] >= -target {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}
	// 没有能从其他步跳过来
	if dp[n-1] < 0 {
		return -1
	}
	return dp[n-1]
}
