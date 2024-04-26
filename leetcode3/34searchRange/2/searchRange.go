package main

import (
	"fmt"
)

func main() {
	num := []int{5, 7, 7, 8, 8, 8, 10}
	// fmt.Println(searchRight(num, 7))
	fmt.Println(searchLeft(num, 7))
	// fmt.Println(searchRange(num, 7))
}

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

// 找左边界 x>=target 的所有区间
func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	// left一定得是 < right，不能是 <=, 不然会进入死循环
	for left < right {
		mid := left + (right-left)>>1
		// mid向下取整，所以 mid 不会越界，
		// left= mid+1，所以mid 不会小于0
		// 但是检查一下，没有坏处
		if mid < len(nums) && mid >= 0 && nums[mid] >= target {
			// if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	// 检测,所有的都检测一次，是一个好习惯
	if left >= len(nums) || left < 0 || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left+1)>>1
		// mid向上取整，可能会导致 mid 越界
		// 又因为left=mid,使用left 不会减少到0以下
		if mid < len(nums) && mid >= 0 && nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 检测
	if left >= len(nums) || left < 0 || nums[left] != target {
		return -1
	}
	return left
}
