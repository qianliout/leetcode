package main

import (
	"fmt"
)

func main() {
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 10, 11}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}

func firstMissingPositive(nums []int) int {

	// 这种写法会出错，原因是在第二个 for 循环中，nums[i]的值有变动
	// for i, ch := range nums {
	// 	for ch > 0 && ch <= len(nums) && ch != nums[ch-1] {
	// 		nums[i], nums[ch-1] = nums[ch-1], nums[i]
	// 	}
	// }
	for i := range nums {
		for (nums[i] > 0 && nums[i] <= len(nums)) && (nums[i] != nums[nums[i]-1]) {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, ch := range nums {
		if ch != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
