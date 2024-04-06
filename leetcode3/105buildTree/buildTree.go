package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := find(inorder, preorder[0])

	root.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	root.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return root
}

func find(nums []int, val int) int {
	for i := range nums {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
