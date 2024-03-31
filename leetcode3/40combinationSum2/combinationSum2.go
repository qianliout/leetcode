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
	dfs(candidates, path, target, 0, used, &ans)
	return ans
}

func dfs(candidates []int, path []int, target, idx int, used map[int]bool, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if target < 0 {
		return
	}
	for i := idx; i < len(candidates); i++ {
		if used[i] {
			continue
		}
		if i > idx && candidates[i] == candidates[i-1] && !used[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target = target - candidates[i]

		dfs(candidates, path, target, i+1, used, ans)

		path = path[:len(path)-1]
		target = target + candidates[i]
	}
}
