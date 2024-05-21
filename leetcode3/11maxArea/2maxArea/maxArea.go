package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	ans, n := 0, len(height)
	if n <= 1 {
		return ans
	}
	l, r := 0, n-1
	lm, rm := height[l], height[r]

	for l < r {
		lm = Max(lm, height[l])
		rm = Max(rm, height[r])
		ans = Max(ans, Min(lm, rm)*(r-l))

		if lm >= rm {
			r--
		} else {
			l++
		}
	}
	return ans
}
