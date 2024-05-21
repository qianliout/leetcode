package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	node := sortedArrayToBST([]int{1, 2, 3, 4, 5})
	InOrderTraversal(node)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums) / 2

	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])
	root := &TreeNode{Val: nums[mid]}
	root.Left = left
	root.Right = right
	return root
}
