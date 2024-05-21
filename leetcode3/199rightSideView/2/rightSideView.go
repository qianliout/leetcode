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
		lev := make([]*TreeNode, 0)
		value := make([]int, 0)
		for _, node := range queue {
			value = append(value, node.Val)
			if node.Left != nil {
				lev = append(lev, node.Left)
			}
			if node.Right != nil {
				lev = append(lev, node.Right)
			}
		}
		ans = append(ans, value[len(value)-1])
		queue = lev
	}
	return ans
}
