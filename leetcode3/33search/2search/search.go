package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{3, 1}, 3))
	fmt.Println(search([]int{3, 5, 1}, 5))
}

func search(nums []int, target int) int {
	// 先找旋转点
	le, ri := 0, len(nums)
	for le < ri {
		mid := le + (ri-le+1)>>1
		// 找大于等nums[0] 的右边界
		// 这里找的是
		if mid >= 0 && mid < len(nums) && nums[mid] >= nums[0] {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	// 根据在左边还是右边，重新确定 left,right 的边界
	if target >= nums[0] {
		// ri此时表示 大于等于nums[0]的右边界，就是大于等于 target 的右边界
		le = 0
	} else {
		le = le + 1
		ri = len(nums)
	}

	for le < ri {
		mid := le + (ri-le+1)>>1
		// 找大于等于 target 的右边界解法
		if mid >= 0 && mid < len(nums) && nums[mid] <= target {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	// 检测
	if le < 0 || le >= len(nums) || nums[le] != target {
		return -1
	}
	return le
}
