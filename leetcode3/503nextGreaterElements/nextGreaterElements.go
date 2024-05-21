package main

import (
	"fmt"
)

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{2}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 5}))
}

// 下一个更大
func nextGreaterElements(nums []int) []int {
	stark := make([]int, 0)
	n := len(nums)
	ans := make([]int, len(nums))
	for i := n*2 - 1; i >= 0; i-- {
		// 递减
		idx := i % n
		for len(stark) > 0 && (stark[len(stark)-1]) <= nums[idx] {
			stark = stark[:len(stark)-1]
		}
		if i <= n-1 {
			if len(stark) == 0 {
				ans[i] = -1
			} else {
				ans[i] = stark[len(stark)-1]
			}
		}
		stark = append(stark, nums[idx])
	}
	return ans
}
