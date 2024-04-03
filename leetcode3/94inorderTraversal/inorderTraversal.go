package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	dfs(root.Left, ans)
	dfs(root.Right, ans)
}
