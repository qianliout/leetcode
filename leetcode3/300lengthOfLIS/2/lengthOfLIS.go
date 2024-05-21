package main

import (
	"math"
)

func main() {

}

func lengthOfLIS(nums []int) int {
	ans := 0
	dp := make([]int, len(nums)+1)
	for i := range nums {
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

func Max(nums ...int) int {
	data := math.MinInt64
	for _, num := range nums {
		if num > data {
			data = num
		}
	}
	return data
}
