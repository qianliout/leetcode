package main

import (
	"fmt"
)

func main() {
	fmt.Println(countRangeSum([]int{-1, 1}, 0, 0))
}

func countRangeSum(nums []int, lower int, upper int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		for j := i; j < len(nums); j++ {
			if i != j {
				sum += nums[j]
			}
			if lower <= sum && upper >= sum {
				ans++
			}
		}
	}
	return ans
}
