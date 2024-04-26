package main

import (
	"fmt"
)

func main() {
	// fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 15))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchLeft([]int{5, 7, 7, 8, 8, 10}, 8))
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// mid := left + (right-left+1)>>1
		mid := (right + left) / 2
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

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)

	return []int{left, right}
}

// 左边界
func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left+1)>>2
		// mid := (right + left) / 2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if left < 0 || left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left+1)/2
		// mid := (right + left) / 2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}

	if right < 0 || right >= len(nums) || nums[right] != target {
		return -1
	}
	return right
}
