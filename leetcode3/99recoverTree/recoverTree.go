package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	exit := make(map[int]*TreeNode)
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	fmt.Println(inOrder(root, exit))
	recoverTree1(root)
	fmt.Println(inOrder(root, exit))
}

func recoverTree1(root *TreeNode) {
	exit := make(map[int]*TreeNode)
	ans := inOrder(root, exit)
	var (
		first *TreeNode
		sec   *TreeNode
	)
	for i := range ans {
		// first 取第一个
		if first == nil && i+1 <= len(ans)-1 && ans[i] > ans[i+1] {
			first = exit[ans[i]]
		}
		// sec 取最后一个
		if i-1 >= 0 && ans[i] < ans[i-1] {
			sec = exit[ans[i]]
		}
	}
	if first != nil && sec != nil {
		first.Val, sec.Val = sec.Val, first.Val
	}
}

func inOrder(root *TreeNode, exit map[int]*TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	exit[root.Val] = root
	left := inOrder(root.Left, exit)
	right := inOrder(root.Right, exit)
	ans = append(ans, left...)
	ans = append(ans, root.Val)
	ans = append(ans, right...)
	return ans
}
