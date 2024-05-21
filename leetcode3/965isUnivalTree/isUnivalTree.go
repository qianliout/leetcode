package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	va := root.Val
	if root.Left != nil && root.Left.Val != va {
		return false
	}
	if root.Right != nil && root.Right.Val != va {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
