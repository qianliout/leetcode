package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(inorder) == 0 {
		return nil
	}

	va := preorder[0]
	idx := find(inorder, va)
	root := &TreeNode{Val: va}
	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func find(inorder []int, va int) int {
	for i, ch := range inorder {
		if ch == va {
			return i
		}
	}
	return 0
}
