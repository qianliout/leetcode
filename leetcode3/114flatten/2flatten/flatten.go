package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

// 这个方法只能转化为链表，但是不满足这个条件
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	if root.Right == nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		return
	}

	flatten(root.Left)
	flatten(root.Right)
	right := root.Right
	for right != nil && right.Right != nil {
		right = right.Right
	}
	right.Right = root.Left
	root.Left = nil
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil {
		flatten(root.Right)
		return
	}
	if root.Right == nil {
		flatten(root.Left)
		root.Right = root.Left
		root.Left = nil
		return
	}

	flatten(root.Left)

	flatten(root.Right)

	right := root.Right
	left := root.Left
	root.Right = left
	for left != nil && left.Right != nil {
		left = left.Right
	}

	left.Right = right

	root.Left = nil
}
