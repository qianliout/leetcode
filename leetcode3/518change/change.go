package main

import (
	"fmt"
)

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))

}

func change(amount int, coins []int) int {
	// dp[i][j] = m
	// 选 i 之前的钱(不包括 i)，凑成总数是 j 元，共有 m 做法
	dp := make([]map[int]int, len(coins)+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	// 初值的定义是一个难点
	// 这里可以理解，没有可选的数，凑成0元，也只有一种做法
	// 第二种方式，用1，0去尝试，会发现如果定义成0的话，后面所的值都是0
	dp[0][0] = 1
	for i := 1; i <= len(coins); i++ {
		ch := coins[i-1]
		for j := 0; j <= amount; j++ {
			no := dp[i-1][j]
			yes := 0

			for k := 1; ch*k <= j; k++ {
				yes += dp[i-1][j-ch*k]
			}

			dp[i][j] = yes + no
		}
	}

	return dp[len(coins)][amount]
}
