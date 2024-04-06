package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root *TreeNode, maxNode, minNode *TreeNode) bool {
	if root == nil {
		return true
	}
	if maxNode != nil && root.Val >= maxNode.Val {
		return false
	}
	if minNode != nil && root.Val <= minNode.Val {
		return false
	}

	return dfs(root.Left, root, minNode) && dfs(root.Right, maxNode, root)
}
