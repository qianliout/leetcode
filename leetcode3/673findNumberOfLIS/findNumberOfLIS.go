package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
}

func findNumberOfLIS(nums []int) int {
	dp1 := make([]int, len(nums)+1) // 以nums[i]结尾的数据的子串长度
	dp2 := make([]int, len(nums)+1) // 考虑以 nums[i]结尾的最长上升子序列的个数。
	maxLen := 0                     // 最长子串的长度
	for i := 0; i < len(nums); i++ {
		dp1[i] = Max(dp1[i], 1)
		dp2[i] = Max(dp2[i], 1)
		for j := i - 1; j >= 0; j-- {
			if nums[i] < nums[j] {
				continue
			}
			dp1[i] = Max(dp1[i], dp1[j]+1)

			if dp1[i] == dp1[j]+1 {
				dp2[i]++
			} else {
				dp2[i] = dp1[j]
			}

			maxLen = Max(maxLen, dp1[i])
		}
	}
	ans := 0
	for i := len(dp2) - 1; i >= 0; i-- {
		if dp1[i] == maxLen {
			ans += dp2[i]
		}
	}
	return ans
}
