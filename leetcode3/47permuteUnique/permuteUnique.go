package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{3, 3, 0, 3}))
}

/*
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	path := make([]int, 0)
	ans := make([][]int, 0)
	used := make(map[int]bool)
	dfs(nums, path, used, &ans)
	return ans
}

func dfs(nums []int, path []int, used map[int]bool, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}

		used[i] = true
		path = append(path, nums[i])
		dfs(nums, path, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}
