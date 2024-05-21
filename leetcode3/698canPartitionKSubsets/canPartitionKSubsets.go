package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
	fmt.Println(canPartitionKSubsets([]int{1, 2, 3, 4}, 3))
	fmt.Println(canPartitionKSubsets([]int{2, 2, 2, 2, 3, 4, 5}, 4))
	fmt.Println(canPartitionKSubsets([]int{10, 1, 10, 9, 6, 1, 9, 5, 9, 10, 7, 8, 5, 2, 10, 8}, 11))
}

func canPartitionKSubsets(nums []int, k int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm%k != 0 {
		return false
	}

	target := sm / k
	mem := make(map[int]bool)
	sort.Ints(nums)
	return dfs(nums, k, target, len(nums)-1, 0, 0, mem)
}

func dfs(nums []int, k, target int, start int, cur int, cnt int, visit map[int]bool) bool {
	if cnt == k {
		return true
	}
	if cur == target {
		return dfs(nums, k, target, len(nums)-1, 0, cnt+1, visit)
	}

	for i := start; i >= 0; i-- {
		if visit[i] || nums[i]+cur > target {
			continue
		}

		visit[i] = true

		if dfs(nums, k, target, i-1, cur+nums[i], cnt, visit) {
			return true
		}

		visit[i] = false
		if cur == 0 {
			return false
		}
	}

	return false
}

func canPartitionKSubsets2(nums []int, k int) bool {
	sm := 0
	for _, ch := range nums {
		sm += ch
	}
	if sm%k != 0 {
		return false
	}

	target := sm / k

	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1

	for i := 1; i <= len(nums); i++ {
		ch := nums[i-1]
		for j := 0; j <= target; j++ {
			no := dp[i-1][j]
			yes := 0
			if j >= ch {
				yes = dp[i-1][j-ch]
			}
			dp[i][j] = no + yes
		}
	}

	return dp[len(nums)][target] >= k
}
