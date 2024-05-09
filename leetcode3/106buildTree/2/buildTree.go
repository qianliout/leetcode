package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) || len(inorder) == 0 {
		return nil
	}
	n := len(postorder)
	va := postorder[len(postorder)-1]
	root := &TreeNode{Val: va}

	idx := find(inorder, va)
	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:n-1])
	return root
}

// 有多少个元素
func find(inorder []int, va int) int {
	for i, ch := range inorder {
		if ch == va {
			return i
		}
	}
	return 0
}
