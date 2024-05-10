package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + dfs(root, targetSum)
}

func dfs(root *TreeNode, target int) int {
	res := 0
	if root == nil {
		return res
	}
	target = target - root.Val
	if target == 0 {
		res++
	}
	res += dfs(root.Left, target)
	res += dfs(root.Right, target)
	return res
}
