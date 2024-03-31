package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make(map[int]bool)
	// dfs(candidates, path, target, 0, &ans)
	dfs2(candidates, path, target, used, 0, &ans)
	return ans
}

func dfs(candidates []int, path []int, target, idx int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if target < 0 {
		return
	}
	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target = target - candidates[i]

		dfs(candidates, path, target, i+1, ans)

		path = path[:len(path)-1]
		target = target + candidates[i]
	}
}

func dfs2(candidates []int, path []int, target int, used map[int]bool, start int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if target < 0 {
		return
	}
	for i := start; i < len(candidates); i++ {
		if used[i] {
			continue
		}

		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		used[i] = true
		path = append(path, candidates[i])
		target = target - candidates[i]

		dfs2(candidates, path, target, used, start+1, ans)

		path = path[:len(path)-1]
		target = target + candidates[i]
		used[i] = false
	}
}
