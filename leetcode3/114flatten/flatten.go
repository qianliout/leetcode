package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	node := &TreeNode{Val: 1}
	node.Left = &TreeNode{Val: 2}
	node.Left.Left = &TreeNode{Val: 3}
	node.Left.Right = &TreeNode{Val: 4}
	node.Right = &TreeNode{Val: 5}
	node.Right.Right = &TreeNode{Val: 6}

	flatten(node)
	PreOrderTraversal(node)
}

func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	left := root.Left
	right := root.Right
	flatten(left)
	flatten(right)
	root.Right = left
	root.Left = nil

	cur := root
	for cur != nil && cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = right
}
