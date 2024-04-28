package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Left != nil && root.Right == nil {
			return root.Left
		}
		if root.Left == nil && root.Right != nil {
			return root.Right
		}
		if root.Left != nil && root.Right != nil {
			left := root.Left
			right := root.Right
			cur := right
			for cur != nil && cur.Left != nil {
				cur = cur.Left
			}
			// 此时 cur 就是root右子树上的最小值
			cur.Left = left
			return right
		}
	}
	return root
}
