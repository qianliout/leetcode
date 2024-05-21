package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs(nums, path, 0, &ans)
	ans = append(ans, []int{})
	return ans
}

func dfs(nums, path []int, start int, ans *[][]int) {
	if len(path) > 0 {
		*ans = append(*ans, append([]int{}, path...))
	}
	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, path, i+1, ans)
		path = path[:len(path)-1]
	}
}
