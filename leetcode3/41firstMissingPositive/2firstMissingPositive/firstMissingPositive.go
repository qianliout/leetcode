package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{1, 2, 3}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		if nums[i] > len(nums) || nums[i] <= 0 {
			continue
		}
		for nums[i]-1 >= 0 && nums[i]-1 < len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
