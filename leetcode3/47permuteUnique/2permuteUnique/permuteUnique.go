package main

import (
	"sort"
)

func main() {

}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	dfs(nums, used, path, &ans)
	return ans
}

func dfs(nums []int, used []bool, path []int, ans *[][]int) {
	if len(path) == len(nums) {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	for i := range nums {
		if used[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && used[i-1] {
			continue
		}
		used[i] = true
		path = append(path, nums[i])
		dfs(nums, used, path, ans)
		used[i] = false
		path = path[:len(path)-1]
	}
}
