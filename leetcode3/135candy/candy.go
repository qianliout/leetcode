package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}

func candy(ratings []int) int {
	left, right := make([]int, len(ratings)), make([]int, len(ratings))

	for i := 0; i < len(ratings); i++ {
		left[i], right[i] = 1, 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	ans := 0
	for i := 0; i < len(ratings); i++ {
		ans = ans + Max(left[i], right[i])
	}

	return ans
}
