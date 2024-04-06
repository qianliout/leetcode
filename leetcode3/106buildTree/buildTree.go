package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	idx := find(inorder, postorder[len(postorder)-1])

	root.Left = buildTree(inorder[:idx], postorder[:idx])
	root.Right = buildTree(inorder[idx+1:], postorder[idx:len(postorder)-1])
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
