package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		queue = lev
	}

	return ans
}
