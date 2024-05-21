package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right == nil {
		return left
	} else if right != nil && left == nil {
		return right
	} else if left != nil && right != nil {
		return root
	}
	return nil
}
