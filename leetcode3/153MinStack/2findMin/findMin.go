package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{2, 1}))
}

// 数组中没有重复值，可以理解成找 小于等于最后一个元素的左边界
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	n := len(nums)
	le, ri := 0, len(nums)
	for le < ri {
		mid := le + (ri-le)>>1
		if mid >= 0 && mid < len(nums) && nums[mid] <= nums[n-1] {
			ri = mid
		} else {
			le = mid + 1
		}
	}
	// if le < 0 || le >= len(nums) {
	// 	return nums[0]
	// }
	return nums[le]
}
