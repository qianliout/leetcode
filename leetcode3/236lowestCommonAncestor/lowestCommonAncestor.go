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

	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}

	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}
