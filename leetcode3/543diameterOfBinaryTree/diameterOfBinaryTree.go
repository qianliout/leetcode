package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	fmt.Println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	dfs(root, &ans)
	return ans - 1
}

// 表示从 root 做为起点最长的非空节点数
func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left, ans)
	right := dfs(root.Right, ans)

	*ans = max(*ans, 1, right+1, left+1, left+right+1)
	return max(right+1, left+1)
}
