package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := range nums {
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}
