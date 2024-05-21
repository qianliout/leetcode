package main

import (
	"fmt"
)

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums)+1)
	right := make([]int, len(nums)+1)
	left[0] = 1
	right[len(right)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] * nums[i]
	}
	ans := make([]int, len(nums))

	for i := range nums {
		ans[i] = left[i] * right[i+1]
	}
	return ans
}
