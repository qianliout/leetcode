package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}

		if len(lev) == 0 {
			return queue[0].Val
		}
		queue = lev
	}

	return 0
}

func dfs(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && mem[root.Left] >= mem[root.Right] {
		return dfs(root.Left, mem)
	} else if root.Right != nil && mem[root.Right] > mem[root.Left] {
		return dfs(root.Right, mem)
	} else {
		return root.Val
	}
}

func findDep(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		mem[root] = 0
		return 0
	}
	if va, ok := mem[root]; ok {
		return va
	}
	le := findDep(root.Left, mem)
	ri := findDep(root.Right, mem)
	mem[root] = max(le, ri) + 1
	return mem[root]
}
