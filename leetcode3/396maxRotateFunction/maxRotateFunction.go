package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}

// timeout
func maxRotateFunction(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	n := len(nums)
	ans := math.MinInt64

	for n > 0 {
		ans = Max(F(nums), ans)
		n--
		fir := nums[0]
		nums = append(nums, fir)
		nums = nums[1:]
	}

	return ans
}

func F(nums []int) int {
	sum := 0
	for i, ch := range nums {
		sum = sum + ch*i
	}
	return sum
}
