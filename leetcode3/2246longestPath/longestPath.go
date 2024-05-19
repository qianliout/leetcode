package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPath([]int{-1, 0, 0, 1, 1, 2}, "abacbe"))
}

func longestPath(parent []int, s string) int {
	n := len(parent)
	// 建图:其中 parent[i] 是节点 i 的父节点
	// 那么，parent[0]就是根节点，他没有父节点，所以值是-1
	node := make([][]int, len(parent))
	for i := 1; i < n; i++ {
		// i 是parent[i]的子节点
		node[parent[i]] = append(node[parent[i]], i)
	}

	ans := 0
	dfs(s, 0, node, &ans)
	// 点的数量是边的数量加1
	return ans + 1
}

// 表示以 x 做为父节点，最大的边长
func dfs(s string, x int, node [][]int, ans *int) int {

	pl := 0 // x做为父节点时的最大长度
	for _, chi := range node[x] {
		// chi 表示 x 的子节点
		chl := dfs(s, chi, node, ans) // 子节点的长度
		// 相连接的字符不能相同
		if s[x] != s[chi] {
			*ans = max(*ans, pl+chl+1)
			pl = max(pl, chl+1)
		}
	}
	return pl
}
