package main

import (
	"fmt"
)

func main() {
	fmt.Println(countSubArray([]int{1, 2, 3}))
}

func countSubArray(nums []int) int {
	ans := 1
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			cnt++
		} else {
			cnt = 0
		}
		ans += cnt
	}
	return ans
}
