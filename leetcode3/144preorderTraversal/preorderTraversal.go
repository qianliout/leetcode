package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func preorderTraversal1(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	left := preorderTraversal1(root.Left)
	right := preorderTraversal1(root.Right)
	ans = append(ans, left...)
	ans = append(ans, right...)
	return ans
}

func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		last := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		ans = append(ans, last.Val)
		if last.Right != nil {
			queue = append(queue, last.Right)
		}
		if last.Left != nil {
			queue = append(queue, last.Left)
		}
	}
	return ans
}
