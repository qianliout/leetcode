package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/treenode"
	. "outback/geeke/leetcode/common/utils"
)

func main() {
	root := &TreeNode{Val: -3}
	fmt.Println(maxPathSum(root))

}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, ans)
	right := dfs(root.Right, ans)
	*ans = Max(*ans, root.Val, root.Val+left, root.Val+right, root.Val+left+right)
	return Max(root.Val, root.Val+left, root.Val+right)
}
