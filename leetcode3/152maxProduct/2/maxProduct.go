package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp1 := make([]int, len(nums))
	dp2 := make([]int, len(nums))
	dp1[0], dp2[0] = nums[0], nums[0]
	ans := dp1[0]
	for i := 1; i < len(nums); i++ {
		ch := nums[i]
		dp2[i] = Min(dp1[i-1]*ch, ch, dp2[i-1]*ch)
		dp1[i] = Max(dp1[i-1]*ch, ch, dp2[i-1]*ch)

		ans = Max(ans, dp1[i])
	}
	return ans
}
