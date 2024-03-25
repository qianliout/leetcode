package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{1, 2, 3}, 4))
}

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
