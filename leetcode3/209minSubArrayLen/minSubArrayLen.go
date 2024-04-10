package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	// fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	// fmt.Println(minSubArrayLen(7, []int{4, 3}))
	// fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	// fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(6, []int{10, 2, 3}))
}

// 可以得出正确结果，但是会超时
func minSubArrayLen2(target int, nums []int) int {
	// dp[i]表示 选定了 nums[i]之后的最小值
	dp := make([]int, len(nums))
	ans := math.MaxInt64
	if len(nums) == 0 {
		return ans
	}
	for i := 0; i < len(nums); i++ {
		th := 0
		for j := i; j >= 0; j-- {
			th += nums[j]
			if th >= target {
				dp[i] = i - j + 1
				ans = Min(ans, dp[i])
				break
			}
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 双指针
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt64
	win := 0
	left, right := 0, 0
	for left <= right && right < len(nums) {
		win += nums[right]
		for left <= right && win >= target {
			ans = Min(ans, right-left+1)
			win -= nums[left]
			left++
		}
		right++
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
