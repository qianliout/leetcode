package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{1, 2, 3}, 4))
}

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的
*/
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	path := make([]int, 0)
	dsf(candidates, path, 0, target, &ans)
	return ans
}

func dsf(candidates []int, path []int, idx, target int, ans *[][]int) {
	if target == 0 {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if target < 0 {
		return
	}
	for i := idx; i < len(candidates); i++ {
		path = append(path, candidates[i])
		target = target - candidates[i]

		dsf(candidates, path, i, target, ans)

		path = path[:len(path)-1]
		target = target + candidates[i]
	}
}
