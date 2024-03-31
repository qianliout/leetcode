package main

import (
	"fmt"
)

func main() {
	fmt.Println(combine(4, 2))
}

/*
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
*/
func combine(n int, k int) [][]int {
	path := make([]int, 0)
	ans := make([][]int, 0)
	dfs(n, k, 1, path, &ans)
	return ans
}

func dfs(n, k, start int, path []int, ans *[][]int) {
	if len(path) == k {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if len(path) > k {
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(n, k, i+1, path, ans)
		path = path[:len(path)-1]
	}
}
