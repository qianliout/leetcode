package main

import (
	"fmt"
)

func main() {
	// fmt.Println(findSubsequences([]int{4, 6, 7, 7}))
	fmt.Println(findSubsequences([]int{4, 4, 3, 2, 1}))
}

func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	dfs(nums, 0, []int{}, &ans)
	return ans
}

func dfs(nums []int, start int, path []int, ans *[][]int) {
	if len(path) >= 2 {
		*ans = append(*ans, append([]int{}, path...))
	}
	used := make(map[int]bool)
	for i := start; i < len(nums); i++ {
		// 同一父节点下，本层重复使用
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			dfs(nums, i+1, path, ans)
			path = path[:len(path)-1]
		}
	}
}
