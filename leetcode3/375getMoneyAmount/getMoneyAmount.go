package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(getMoneyAmount(10))
}

func getMoneyAmount2(n int) int {
	mem := make(map[[2]int]int)
	return dfs(1, n, mem)
}

// 表示从 start-->end 这个数据范围中，要赢得游戏需要的钱数
func dfs(start, end int, mem map[[2]int]int) int {
	if start >= end {
		return 0
	}
	if va, ok := mem[[2]int{start, end}]; ok {
		return va
	}
	ans := math.MaxInt64
	for i := start; i <= end; i++ {
		// 要保证一定能猜测到，就只能选择两边的较大值
		cur := Max(dfs(start, i-1, mem), dfs(i+1, end, mem)) + i
		// 在所有我们可以决策的数值之间取最优
		ans = Min(ans, cur)
	}
	mem[[2]int{start, end}] = ans
	return ans
}

func getMoneyAmount(n int) int {
	// dp[i][j] 表示从 start-->end 这个数据范围中，要赢得游戏需要的钱数
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 先确定右边,再从右边扩散到左边
	// 因为左边构建好了，再能计算后面的
	for r := 1; r <= n; r++ {
		for l := r - 1; l >= 1; l-- {
			ans := math.MaxInt64
			for k := l; k < r; k++ {
				ans = Min(ans, Max(dp[l][k-1], dp[k+1][r])+k)
			}
			dp[l][r] = ans
		}
	}
	return dp[1][n]
}
