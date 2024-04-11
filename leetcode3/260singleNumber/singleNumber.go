package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
	fmt.Println(singleNumber([]int{2, 1, 2, 3, 4, 1}))
}

func singleNumber(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	first := 0
	for i := 0; i < len(nums); i++ {
		first = first ^ nums[i]
	}
	// 最后一个1
	diff := first & -first
	left, right := 0, 0
	for _, ch := range nums {
		if ch&diff == 0 {
			left = left ^ ch
		} else {
			right = right ^ ch
		}
	}
	return []int{left, right}
}
