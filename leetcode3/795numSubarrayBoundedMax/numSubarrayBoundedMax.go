package main

import (
	"fmt"
)

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
}

/*
给定一个元素都是正整数的数组 A ，正整数 L  以及  R (L <= R)。
求连续、非空且其中最大元素满足大于等于 L  小于等于 R 的子数组个数。
*/
func numSubarrayBoundedMax(nums []int, left int, right int) int {
	// return notGreater(nums, right) - notGreater(nums, left-1)
	return slidingWindow(nums, right) - slidingWindow(nums, left-1)
}

// 最大元素是 k 的子串个数
func notGreater(nums []int, k int) int {
	ans, cnt := 0, 0
	// 以nums[i]结尾的子数组的长度
	dp := make([]int, len(nums))

	for i, ch := range nums {
		if ch <= k {
			cnt++
		} else {
			// 为啥是0呢，因为 如果 ch>k,那么以 ch结尾的数组都不可取，所以是0
			cnt = 0
		}
		dp[i] = cnt
	}
	for _, ch := range dp {
		ans += ch
	}

	return ans
}

func slidingWindow(nums []int, k int) int {
	// cnt相当于滑动窗口
	ans, cnt := 0, 0
	left, right := 0, 0

	for left <= right && right < len(nums) {
		if nums[right] > k {
			left = right
			cnt = 0
		} else {
			cnt++
		}
		right++
		ans += cnt
	}
	return ans
}

// 最容易理解的方法
/**
func findSubstringInWraproundString(s string) int {
	pre := 0
	// 把a-z映射到 0-25
	// dp[i]以字母i+'a'结尾，存在的最长的子串的长度
	dp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if i > 0 && check(s[i-1], s[i]) {
			pre += 1
		} else {
			pre = 1
		}
		idx := int(s[i] - 'a')
		dp[idx] = Max(dp[idx], pre)
	}
	ans := 0
	// 前缀和的思想
	for _, ch := range dp {
		ans += ch
	}
	return ans
}

*/
