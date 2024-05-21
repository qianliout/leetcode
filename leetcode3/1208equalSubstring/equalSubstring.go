package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func equalSubstring2(s string, t string, maxCost int) int {
	left, right := 0, 0
	ans := 0

	for left <= right && right < len(s) {
		maxCost -= AbsSub(int(s[right]), int(t[right]))
		right++
		for maxCost < 0 {
			maxCost += AbsSub(int(s[left]), int(t[left]))
			left++
		}
		ans = Max(ans, right-left)
	}
	return ans
}

func equalSubstring(s string, t string, maxCost int) int {
	sum := make([]int, len(t)+1)
	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + AbsSub(int(t[i-1]), int(s[i-1]))
	}
	left, right := 1, len(sum)
	for left < right {
		// 右端点
		mid := left + (right-left+1)>>1
		if check(sum, mid, maxCost) {
			left = mid
		} else {
			right = mid - 1

		}
	}

	if check(sum, left, maxCost) {
		return left
	}
	return 0
}
func check(sum []int, mid int, maxCost int) bool {
	for i := mid; i < len(sum); i++ {
		if sum[i]-sum[i-mid] <= maxCost {
			return true
		}
	}

	return false
}
