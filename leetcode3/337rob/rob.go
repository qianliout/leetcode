package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func rob(root *TreeNode) int {
	mem := make(map[*TreeNode]int)
	return dfs(root, mem)
}

// 可以得到结果，但是会timeout
func rob2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	ans1 := root.Val
	if root.Left != nil {
		ans1 += rob(root.Left.Left)
		ans1 += rob(root.Left.Right)
	}
	if root.Right != nil {
		ans1 += rob(root.Right.Left)
		ans1 += rob(root.Right.Right)
	}
	ans2 := rob(root.Left) + rob(root.Right)

	if ans1 > ans2 {
		return ans1
	}
	return ans2
}

func dfs(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}

	ans1 := root.Val
	if root.Left != nil {
		ans1 += dfs(root.Left.Left, mem)
		ans1 += dfs(root.Left.Right, mem)
	}
	if root.Right != nil {
		ans1 += dfs(root.Right.Left, mem)
		ans1 += dfs(root.Right.Right, mem)
	}

	ans2 := dfs(root.Left, mem) + dfs(root.Right, mem)

	if ans1 > ans2 {
		mem[root] = ans1
		return ans1
	}
	mem[root] = ans2
	return ans2
}
