package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func postorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	left := postorderTraversal2(root.Left)
	right := postorderTraversal2(root.Right)
	ans = append(ans, left...)
	ans = append(ans, right...)
	ans = append(ans, root.Val)
	return ans
}

func postorderTraversal(root *TreeNode) []int {
	ans, stark1, stark2 := make([]int, 0), make([]*TreeNode, 0), make([]*TreeNode, 0)
	if root == nil {
		return ans
	}
	stark1 = append(stark1, root)
	for len(stark1) > 0 {
		last := stark1[len(stark1)-1]
		stark1 = stark1[:len(stark1)-1]
		stark2 = append(stark2, last)
		if last.Left != nil {
			stark1 = append(stark1, last.Left)
		}
		if last.Right != nil {
			stark1 = append(stark1, last.Right)
		}
	}
	for i := len(stark2) - 1; i >= 0; i-- {
		ans = append(ans, stark2[i].Val)
	}
	return ans
}
