package main

import (
	. "outback/geeke/leetcode/common/treenode"
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func rob(root *TreeNode) int {
	mem := make(map[*TreeNode]int)
	return dfs(root, mem)
}

func dfs(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}

	a := root.Val
	if root.Left != nil {
		a += dfs(root.Left.Left, mem) + dfs(root.Left.Right, mem)
	}
	if root.Right != nil {
		a += dfs(root.Right.Left, mem) + dfs(root.Right.Right, mem)
	}

	b := dfs(root.Left, mem) + dfs(root.Right, mem)
	mem[root] = Max(a, b)
	return Max(a, b)
}
