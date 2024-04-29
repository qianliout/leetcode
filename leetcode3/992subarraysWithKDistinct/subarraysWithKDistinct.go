package main

import (
	"fmt"
)

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))

}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMostK(nums, k) - atMostK(nums, k-1)
}

// 不超过K的数的子数据组
func atMostK(nums []int, k int) int {
	window := make(map[int]int)
	left, right := 0, 0
	ans := 0
	// dp[i] 表示以 nums[i]结尾的符合要求的字数组的长度
	dp := make([]int, len(nums))

	for left <= right && right < len(nums) {
		ri := nums[right]
		if window[ri] == 0 {
			k--
		}
		window[ri]++

		for k < 0 {
			le := nums[left]
			window[le]--
			left++

			if window[le] == 0 {
				k++
			}
		}
		dp[right] = right - left + 1
		right++
		// [0,1,2],right-left = 2 不是3个元素，那这里为啥能得出正确答案呢，因为前面 right++了
		// ans += right - left
		// ans = Max(ans, right-left)
	}

	for _, ch := range dp {
		ans += ch
	}
	return ans
}

// 前缀和的思想
func atMostK2(nums []int, k int) int {
	window := make(map[int]int)
	left, right := 0, 0
	ans := 0

	for left <= right && right < len(nums) {
		ri := nums[right]
		if window[ri] == 0 {
			k--
		}
		window[ri]++
		right++

		for k < 0 {
			le := nums[left]
			window[le]--
			left++

			if window[le] == 0 {
				k++
			}
		}
		// [0,1,2],right-left = 2 不是3个元素，那这里为啥能得出正确答案呢，因为前面 right++了
		ans += right - left
	}

	return ans
}
