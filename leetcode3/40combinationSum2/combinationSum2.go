package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	dfs(candidates, path, target, 0, &ans)
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
		if i != idx && candidates[i] == candidates[i-1] {
			continue
		}
		path = append(path, candidates[i])
		target = target - candidates[i]

		dfs(candidates, path, target, i+1, ans)

		path = path[:len(path)-1]
		target = target + candidates[i]
	}
}
