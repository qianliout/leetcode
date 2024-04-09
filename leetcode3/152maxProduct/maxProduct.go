package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxDP := make([]int, len(nums))
	minDP := make([]int, len(nums))
	maxDP[0], minDP[0] = nums[0], nums[0]
	ans := maxDP[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i] = Max(nums[i], maxDP[i-1]*nums[i], minDP[i-1]*nums[i])
		minDP[i] = Min(nums[i], maxDP[i-1]*nums[i], minDP[i-1]*nums[i])
		ans = Max(ans, maxDP[i])
	}

	return ans
}
