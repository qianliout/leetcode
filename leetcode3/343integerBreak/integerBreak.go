package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(integerBreak(10))
	fmt.Println(integerBreak(3))
	fmt.Println(integerBreak(1))
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(6))
}

func integerBreak(n int) int {
	// dp[i] 表示 i 的结果值
	// dp[i]: 正整数i对应的最大乘积
	dp := make(map[int]int)
	dp[0], dp[1], dp[2] = 0, 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i; j++ {
			// 拆分成两个 j*(i-j)
			// 拆分成两个及以上 dp[j]*(i-j)
			dp[i] = Max(dp[i], dp[j]*(i-j), j*(i-j))
		}
	}
	return dp[n]
}
