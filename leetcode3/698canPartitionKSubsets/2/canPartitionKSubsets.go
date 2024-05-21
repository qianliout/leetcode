package main

import (
	"sort"
)

func main() {

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
	sort.Ints(nums)
	mem := make(map[int]bool)
	return dfs(nums, k, target, 0, 0, 0, mem)
}

func dfs(nums []int, k, target int, start int, cur int, cnt int, visit map[int]bool) bool {
	if cnt == k {
		return true
	}
	if cur == target {
		return dfs(nums, k, target, 0, 0, cnt+1, visit)
	}

	for i := start; i < len(nums); i++ {
		if visit[i] || nums[i]+cur > target {
			continue
		}

		visit[i] = true
		if dfs(nums, k, target, i+1, cur+nums[i], cnt, visit) {
			return true
		}
		visit[i] = false
		if cur == 0 {
			return false
		}
	}

	return false
}
