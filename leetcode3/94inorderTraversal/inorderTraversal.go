package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func inorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, ans)
	*ans = append(*ans, root.Val)
	dfs(root.Right, ans)
}

func inorderTraversal3(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	left := inorderTraversal3(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal3(root.Right)
	left = append(left, right...)
	return left
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 || root != nil {
		if root != nil {
			queue = append(queue, root)
			root = root.Left
		} else {
			last := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			ans = append(ans, last.Val)
			root = last.Right
		}
	}
	return ans
}
