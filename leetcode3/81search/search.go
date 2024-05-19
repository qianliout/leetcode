package main

import (
	"fmt"
)

func main() {
	// fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	// fmt.Println(search([]int{3, 4, 1}, 1))
	fmt.Println(search([]int{2, 2, 2, 3, 2, 2, 2}, 3))
}

func search(nums []int, target int) bool {
	le, ri := 0, len(nums)
	// 恢复二段性
	for le < ri && nums[0] == nums[ri-1] {
		ri--
	}

	for le < ri {
		mid := le + (ri-le+1)>>1
		if mid >= 0 && mid < len(nums) && nums[mid] >= nums[0] {
			le = mid
		} else {
			ri = mid - 1
		}
	}

	if target >= nums[0] {
		le = 0
	} else {
		le = le + 1
		ri = len(nums)
	}

	for le < ri {
		mid := le + (ri-le+1)>>1
		// 右边界解法
		if mid >= 0 && mid < len(nums) && nums[mid] <= target {
			le = mid
		} else {
			ri = mid - 1
		}
	}
	return le >= 0 && le < len(nums) && nums[le] == target
}

func find(nums []int, l int, r int, target int) bool {
	n := r
	for l < r {
		mid := l + (r-l)>>1
		// 找左边界
		if mid >= 0 && mid < n && nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l < 0 || r >= n || nums[l] != target {
		return false
	}
	return true
}
