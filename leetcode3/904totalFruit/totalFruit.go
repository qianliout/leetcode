package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(totalFruit([]int{1, 2, 1}))
	fmt.Println(totalFruit([]int{0, 1, 2, 2}))
	fmt.Println(totalFruit([]int{1, 2, 3, 2, 2}))
}

func totalFruit(fruits []int) int {
	return atMostK(fruits, 2)
}

func atMostK(nums []int, k int) int {
	window := make(map[int]int)
	left, right := 0, 0
	ans := 0
	for left <= right && right < len(nums) {
		ri := nums[right]
		right++

		if window[ri] == 0 {
			// 之前没有这个数，这次是新加的，使用了一个篮子
			k--
		}
		window[ri]++

		for left <= right && k < 0 {
			le := nums[left]
			window[le]--
			left++
			if window[le] == 0 {
				k++
			}
		}
		// [0,1,2],right-left = 2 不是3个元素，那这里为啥能得出正确答案呢，因为前面 right++了
		ans = Max(ans, right-left)
	}

	return ans
}
