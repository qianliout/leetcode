package main

import (
	"fmt"
)

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}

func canCross(stones []int) bool {
	if len(stones) <= 1 {
		return true
	}
	if stones[1] != 1 {
		return false
	}
	// dp[i][j]表示上一跳的步长是 j 能不能跳到 idx=i 处
	dp := make([]map[int]bool, len(stones))
	for i := range dp {
		dp[i] = make(map[int]bool)
	}
	// 初值
	dp[1][1] = true
	for i := 2; i < len(stones); i++ {
		for j := 1; j < i; j++ {
			k := stones[i] - stones[j]
			if k <= j+1 {
				dp[i][k] = dp[j][k-1] || dp[j][k] || dp[j][k+1]
			}
		}
	}

	for _, v := range dp[len(stones)-1] {
		if v {
			return true
		}
	}
	return false
}
