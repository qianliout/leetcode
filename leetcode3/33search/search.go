package main

import (
	"fmt"
)

func main() {
	// fmt.Println(search4([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// fmt.Println(search4([]int{3, 1}, 3))
	// fmt.Println(search4([]int{3, 1}, 1))
	fmt.Println(search4([]int{1, 3}, 1))
	// fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 3))
}

// 数组中的值 互不相同 。
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		// 先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段
		if nums[mid] >= nums[left] {
			// 再判断 target 是在 mid 的左边还是右边，从而调整左右边界 left 和 right
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func search3(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		// 当目标在左边部分
		mid := left + (right-left)>>1
		if target >= nums[0] {
			// 当 mid 在左边部分，且值小于 target 时，向右收缩
			// 当 mid 在右边部分，或在左边部分且值 大于等于 target 时，向左收缩
			// 最后结果是第一个大于等于 target 的下标
			if nums[mid] >= nums[0] && nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
			// 当目标在右边部分
		} else {
			if nums[mid] >= nums[0] || nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func search4(nums []int, target int) int {
	left, right := 0, len(nums)
	// 找旋转点
	for left < right {
		// 右边界模板
		mid := left + (right-left+1)>>1
		if mid >= 0 && mid < len(nums) && nums[mid] >= nums[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 根据在左边还是右边，重新确定 left,right 的边界
	if target >= nums[0] {
		left = 0
	} else {
		left = left + 1
		right = len(nums)
	}

	for left < right {
		mid := left + (right-left)>>1
		// 左边界模板
		if mid >= 0 && mid < len(nums) && nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < 0 || left >= len(nums) || nums[left] != target {
		return -1
	}

	return left
}

// 数组中的值互不相同,可以用两次二分
func search2(nums []int, target int) int {
	// 第一次二分，找到旋转点
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] > nums[0] {
			left = mid
		} else {
			right = mid - 1
		}
	}

	// 确定 target 在旋转点的左边还是右边
	if target >= nums[0] {
		left = 0
	} else {
		left = left + 1
		right = len(nums) - 1
	}
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}
