package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
	for len(stark) > 0 {
		lev := make([]int, 0)
		le := make([]*TreeNode, 0)
		for _, no := range stark {
			lev = append(lev, no.Val)
			if no.Left != nil {
				le = append(le, no.Left)
			}
			if no.Right != nil {
				le = append(le, no.Right)
			}
		}
		ans = append(ans, lev)
		stark = le
	}
	return ans
}
