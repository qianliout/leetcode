package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func rangeSumBST(root *TreeNode, low int, high int) int {
	ans := make([]int, 0)
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		lev := make([]*TreeNode, 0)
		for _, no := range queue {
			if no.Val >= low && no.Val <= high {
				ans = append(ans, no.Val)
			}
			if no.Left != nil {
				lev = append(lev, no.Left)
			}
			if no.Right != nil {
				lev = append(lev, no.Right)
			}
		}
		queue = lev
	}
	res := 0
	for _, ch := range ans {
		res += ch
	}
	return res
}
