package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	tre := &TreeNode{Val: 10}
	tre.Left = &TreeNode{Val: 5}
	tre.Left.Left = &TreeNode{Val: 3}
	tre.Left.Left.Left = &TreeNode{Val: 3}
	tre.Left.Left.Right = &TreeNode{Val: -2}
	tre.Left.Right = &TreeNode{Val: 2}
	tre.Left.Right.Right = &TreeNode{Val: 1}
	tre.Right = &TreeNode{Val: -3}
	tre.Right.Right = &TreeNode{Val: 11}

	fmt.Println(pathSum(tre, 8))
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	ans := dfs(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)

	return ans
}

func dfs(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}
	ans := 0
	if root.Val == target {
		ans++
	}
	left := dfs(root.Left, target-root.Val)
	right := dfs(root.Right, target-root.Val)
	return ans + left + right
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := sum(root.Left)
	right := sum(root.Right)
	root.Val = root.Val + left + right
	return root.Val
}
