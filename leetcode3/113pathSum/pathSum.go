package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	dfs(root, targetSum, []int{}, &ans)
	return ans
}

func dfs(root *TreeNode, tag int, path []int, ans *[][]int) {
	if root == nil {
		return
	}
	tag = tag - root.Val
	path = append(path, root.Val)
	if tag == 0 && root.Left == nil && root.Right == nil {
		*ans = append(*ans, append([]int{}, path...))
		return
	}
	if tag < 0 {
		return
	}
	dfs(root.Left, tag, path, ans)
	dfs(root.Right, tag, path, ans)
}
