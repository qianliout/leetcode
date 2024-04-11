package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

// 二叉搜索树
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	left := count(root.Left)
	if left == k-1 {
		return root.Val
	} else if left > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-left-1)
	}
}

func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + count(root.Left) + count(root.Right)
}
