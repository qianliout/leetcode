package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	// 保证能跳到
	for i := 0; i < len(nums); i++ {
		// 初值
		dp[i] = i
		for j := i - 1; j >= 0; j-- {
			if i-j <= nums[j] && dp[i] > dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[len(nums)-1]
}

func jump(nums []int) int {
	cnt, end, maxPos := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		if maxPos < i+nums[i] {
			maxPos = i + nums[i]
		}
		// end表示，循环到 i-1时，最远能走多远
		// maxPos表示循环到 i 时，最远能走多远
		if end == i && i != len(nums)-1 {
			cnt++
			end = maxPos
		}
	}

	return cnt
}
