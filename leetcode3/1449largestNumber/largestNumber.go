package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(largestNumber([]int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9))
}

/*
输入：cost = [4,3,2,5,6,7,2,5,5], target = 9
输出："7772"
*/
func largestNumber1(cost []int, target int) string {
	ans := make([]string, 0)

	// dp[i][j] 表示 前 i 个数，最大消耗是 target 时,组成的最大的数有多少位
	// dp[9][12] = 4 表示使用1-9的这9个数，最大消耗12时，组成的最大的数是4位
	dp := make([][]int, 10)
	for i := range dp {
		dp[i] = make([]int, target+1)
		// 如果不写这样的初值，得不到正确的结果
		for j := range dp[i] {
			dp[i][j] = math.MinInt / 2
		}
	}

	dp[0][0] = 0
	for i := 1; i <= 9; i++ {
		ch := cost[i-1]

		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := math.MinInt / 2

			for k := 1; k*ch <= j; k++ {
				yes = Max(yes, dp[i-1][j-ch*k]+k)
			}

			dp[i][j] = Max(no, yes)
		}
	}
	if dp[9][target] < 0 {
		return "0"
	}

	// 然后再找每一个数，是否可用，从大到小的找可以找到最大值
	for i := 9; i >= 1; i-- {
		j := target
		ch := cost[i-1]

		// 因为这个判断dp[i][j] == dp[i][j-ch]+1，所以初值一定是一个很小的负数
		for j >= ch && dp[i][j] > 0 && dp[i][j] == dp[i][j-ch]+1 {
			ans = append(ans, strconv.Itoa(i))
			j = j - ch
			target = target - ch
		}
	}

	return strings.Join(ans, "")
}

func largestNumber(cost []int, target int) string {
	ans := make([]string, 0)

	dp := make([]int, target+1)
	// 初值是这道题的关键
	for i := range dp {
		dp[i] = math.MinInt32
	}
	dp[0] = 0
	for i := 1; i <= 9; i++ {
		ch := cost[i-1]
		for j := 0; j <= target; j++ {
			no := dp[j]
			yes := math.MinInt32
			if j >= ch {
				yes = dp[j-ch] + 1
			}
			dp[j] = Max(no, yes)
		}
	}
	if dp[target] < 0 {
		return "0"
	}
	for i := 9; i >= 1; i-- {
		ch := cost[i-1]
		for target >= ch && dp[target] > 0 && dp[target-ch]+1 == dp[target] {
			ans = append(ans, strconv.Itoa(i))
			target = target - ch
		}
	}
	return strings.Join(ans, "")
}
