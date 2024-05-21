package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	stark := make([]*TreeNode, 0)
	stark = append(stark, root)
	right := true
	for len(stark) > 0 {
		lev := make([]int, 0)
		stark2 := make([]*TreeNode, 0)
		for _, no := range stark {
			lev = append(lev, no.Val)
			if no.Left != nil {
				stark2 = append(stark2, no.Left)
			}
			if no.Right != nil {
				stark2 = append(stark2, no.Right)
			}
		}

		if !right {
			l, r := 0, len(lev)-1
			for l < r {
				lev[l], lev[r] = lev[r], lev[l]
				l++
				r--
			}
		}
		ans = append(ans, lev)
		right = !right
		stark = stark2
	}
	return ans
}
