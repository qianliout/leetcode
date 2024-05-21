package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeightII([]int{1, 1, 4, 2, 2}))
}

func lastStoneWeightII(stones []int) int {
	n := len(stones)
	sm := 0
	for _, ch := range stones {
		sm += ch
	}
	target := sm >> 1
	// 前i 个元素，成本不超过 j 时的最大价值
	// 成本和价值都是 stones[i]
	dp := make([]int, target+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		ch := stones[i-1]
		for j := target; j >= ch; j-- {
			dp[j] = Max(dp[j], dp[j-ch]+ch)
		}
	}

	return AbsSub(sm-dp[target], dp[target])
}
