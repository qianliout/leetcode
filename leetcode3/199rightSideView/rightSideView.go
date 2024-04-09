package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		if len(lev) > 0 {
			ans = append(ans, lev[len(lev)-1].Val)
		}
		queue = lev
	}

	return ans
}
