package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) != len(preorder) || len(postorder) == 0 {
		return nil
	}

	first := preorder[0]
	root := &TreeNode{Val: first}
	// 未完成
	return root
}

func find(postorder []int, va int) int {
	for i, ch := range postorder {
		if ch == va {
			return i
		}
	}
	return 0
}
