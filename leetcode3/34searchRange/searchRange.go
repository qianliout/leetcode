package main

import (
	"fmt"
)

func main() {

	// num := []int{2, 2, 3, 3, 3, 3}
	num := []int{5, 7, 7, 8, 8, 8, 10}
	fmt.Println(searchLeft(num, 10))
	fmt.Println(searchRight(num, 8))

	fmt.Println(searchRange(num, 8))
}

func searchRange(nums []int, target int) []int {
	left := searchLeft(nums, target)
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	left, rig := 0, len(nums)-1
	for left <= rig {
		mid := left + (rig-left)/2
		if nums[mid] == target {
			rig = mid - 1
		} else if nums[mid] > target {
			rig = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if left >= len(nums) || left < 0 || nums[left] != target {
		return -1
	}
	return left
}

func searchRight(nums []int, target int) int {
	left, rig := 0, len(nums)-1
	for left <= rig {
		mid := left + (rig-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			rig = mid - 1

		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	if rig < 0 || rig >= len(nums) || nums[rig] != target {
		return -1
	}
	return rig
}
