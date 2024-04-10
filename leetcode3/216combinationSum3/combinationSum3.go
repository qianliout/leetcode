package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {

	ans := make([][]int, 0)
	if 9*k < n {
		return ans
	}

	path := make([]int, 0)
	dfs(k, 1, n, path, &ans)
	return ans
}

func dfs(k, start, n int, path []int, ans *[][]int) {
	if n == 0 && len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if n < 0 || len(path) > k {
		return
	}
	for i := start; i <= 9; i++ {
		path = append(path, i)
		dfs(k, i+1, n-i, path, ans)
		path = path[:len(path)-1]
	}
}
