package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/
func permute(nums []int) [][]int {
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

		used[i] = true
		path = append(path, nums[i])
		dfs(nums, path, used, ans)
		path = path[:len(path)-1]
		used[i] = false
	}
}
