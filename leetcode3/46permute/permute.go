package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

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
