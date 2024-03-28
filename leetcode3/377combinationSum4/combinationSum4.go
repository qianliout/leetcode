package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 0))
}

// https://leetcode.cn/problems/combination-sum-iv/solutions/124393/xi-wang-yong-yi-chong-gui-lu-gao-ding-bei-bao-wen-/
func combinationSum4(nums []int, target int) int {
	dp := make(map[int]int)
	// 动态规划的边界是 dp[0]=1\textit{dp}[0]=1dp[0]=1。只有当不选取任何元素时，元素之和才为 000，因此只有 111 种方案。
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := range nums {
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}
