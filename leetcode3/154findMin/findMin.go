package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
}

// 有重复值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	// 假设发生了折叠，那么前半段一定会大于后半段，所以mid 和 right 比较就好
	for left < right {
		if nums[left] == nums[right] {
			left++
			continue
		}
		// 最小值，相当于左边界
		mid := left + (right-left)>>1
		// 在范围中
		if mid >= 0 && mid < len(nums) && nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 退出循环时，left可能是 right+1,对于这个题来说，这也是答案

	// 检查结果,查最小值，这个值一定会有的，
	return nums[left]
}
