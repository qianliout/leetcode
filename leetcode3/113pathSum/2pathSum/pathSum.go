package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	dfs(root, []int{}, targetSum, &res)
	return res
}

func dfs(root *TreeNode, path []int, target int, ans *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	target = target - root.Val

	if root.Left == nil && root.Right == nil {
		if target == 0 {
			*ans = append(*ans, append([]int{}, path...))
		}
		return
	}
	dfs(root.Left, path, target, ans)
	dfs(root.Right, path, target, ans)
	path = path[:len(path)-1]
	target = target + root.Val
}
