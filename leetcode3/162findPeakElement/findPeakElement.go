package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPeakElement([]int{1}))
}

func findPeakElement(nums []int) int {
	nums = append(nums, math.MinInt64)
	nums = append([]int{math.MinInt64}, nums...)

	for i := 1; i < len(nums)-1; i++ {
		if nums[i-1] < nums[i] && nums[i+1] < nums[i] {
			return i - 1
		}
	}
	return 0
}
