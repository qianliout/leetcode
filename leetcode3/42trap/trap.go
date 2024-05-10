package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(nums))
}

func trap2(height []int) int {
	// 递减
	stark := make([]int, 0)
	ans := 0
	for i, ch := range height {
		for len(stark) > 0 && height[stark[len(stark)-1]] < ch {
			last := stark[len(stark)-1]

			stark = stark[:len(stark)-1]
			if len(stark) == 0 {
				break
			}

			l := stark[len(stark)-1]
			r := i
			h := Min(height[l], height[r]) - height[last]
			ans += (r - l - 1) * h
		}

		stark = append(stark, i)
	}
	return ans
}

func trap3(height []int) int {
	n := len(height)
	l, r := 0, n-1
	lm, rm := height[0], height[n-1]

	ans := 0
	for l < r {
		if lm < rm {
			ans += lm - height[l]
			l++
			lm = Max(lm, height[l])
		} else {
			ans += rm - height[r]
			r--
			rm = Max(rm, height[r])
		}
	}
	return ans
}

func trap(height []int) int {

	n := len(height)
	l, r := 0, n-1
	lm, rm := height[0], height[n-1]
	ans := 0
	for l < r {
		lm = Max(lm, height[l])
		rm = Max(rm, height[r])
		if lm < rm {
			ans += lm - height[l]
			l++
		} else {
			ans += rm - height[r]
			r--
		}
	}
	return ans
}
