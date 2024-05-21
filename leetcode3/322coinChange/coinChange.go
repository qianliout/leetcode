package main

import (
	"fmt"
	"math"
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for _, ch := range coins {
			if i-ch >= 0 && dp[i-ch] != math.MaxInt64 {
				dp[i] = Min(dp[i], dp[i-ch]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}
