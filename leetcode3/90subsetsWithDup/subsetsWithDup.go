package main

import (
	"sort"
)

func main() {

}

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的
子集
（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
*/
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0)
	used := make(map[int]bool)
	ans := make([][]int, 0)
	dfs(nums, path, 0, used, &ans)

	ans = append(ans, []int{})
	return ans
}

func dfs(nums []int, path []int, start int, used map[int]bool, ans *[][]int) {
	if len(path) > 0 {
		*ans = append(*ans, append([]int{}, path...))
	}

	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}

		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		path = append(path, nums[i])
		dfs(nums, path, i, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}
