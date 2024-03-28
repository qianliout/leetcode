package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	maxPos := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxPos < i {
			return false
		}
		if maxPos < nums[i]+i {
			maxPos = nums[i] + i
		}
	}
	return true
}
