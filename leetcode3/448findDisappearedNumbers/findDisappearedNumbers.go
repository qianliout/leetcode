package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 {
			if nums[i]-1 < 0 || nums[i]-1 >= len(nums) || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]

		}
	}
	ans := make([]int, 0)
	for i, ch := range nums {
		if i+1 != ch {
			ans = append(ans, i+1)
		}
	}
	return ans
}
