package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		he := Min(height[r], height[l]) * (r - l)
		if he > ans {
			ans = he
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
