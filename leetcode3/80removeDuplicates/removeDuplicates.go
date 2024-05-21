package main

import (
	"fmt"
)

func main() {
	removeDuplicates([]int{0, 0, 0, 0, 1, 1, 1, 1})
	removeDuplicates([]int{1, 1, 1, 2, 2, 3})
}

func removeDuplicates(nums []int) int {
	start := 0
	for i := 0; i < len(nums); i++ {
		if start >= 2 && nums[i] == nums[start-1] && nums[i] == nums[start-2] {
			continue
		}
		nums[start] = nums[i]
		start++
	}
	fmt.Println(nums[:start])
	return start
}
