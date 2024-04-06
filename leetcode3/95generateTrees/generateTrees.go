package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	node := generateTrees(3)
	for _, n := range node {
		PreOrderTraversal(n)
	}
}

func generateTrees(n int) []*TreeNode {
	return dfs(1, n)
}

func dfs(start, end int) []*TreeNode {
	ans := make([]*TreeNode, 0)
	if start <= 0 || end <= 0 || start > end {
		return []*TreeNode{nil}
	}
	for i := start; i <= end; i++ {
		left := dfs(start, i-1)
		right := dfs(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i}
				root.Left = l
				root.Right = r
				ans = append(ans, root)
			}
		}
	}
	return ans
}
