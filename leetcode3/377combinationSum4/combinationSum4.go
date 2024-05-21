package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4D([]int{1, 2, 3}, 4))
}

// https://leetcode.cn/problems/combination-sum-iv/solutions/124393/xi-wang-yong-yi-chong-gui-lu-gao-ding-bei-bao-wen-/
func combinationSum4D(nums []int, target int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for j := range nums {
			dp[i] += dp[i-nums[j]]
		}
	}
	return dp[target]
}

// 请注意，这个方法下 [2,2,1]和[2,1,2]会被看做相同的数列
func combinationSum4C(nums []int, target int) int {
	dp := make([]map[int]int, len(nums)+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := 0
			for k := 1; j >= k*ch; k++ {
				yes = yes + dp[i-1][j-k*ch]
			}
			dp[i][j] = no + yes
		}
	}

	return dp[len(nums)][target]
}

func combinationSum4(nums []int, target int) int {
	// dp[i][j] 组合长度是i,凑成值是 j ，共公有多少种凑法
	// 注意这里的 i 是组合长度，不是 nums 的长度
	// 组合长度最长就是 target (可能的情况下每个都选1),
	dp := make([]map[int]int, target+1)
	for i := range dp {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 1

	ans := 0
	for i := 1; i <= target; i++ {
		for j := 0; j <= target; j++ {
			for k := 1; k <= len(nums); k++ {
				ch := nums[k-1]
				if j < ch {
					break
				}
				dp[i][j] += dp[i-1][j-ch]
			}
		}
		ans += dp[i][target]
	}
	return ans
}
