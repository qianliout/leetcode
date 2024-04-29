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

func atMostK(nums []int, k int) int {
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
		// ans = Max(ans, right-left)
	}

	return ans
}
