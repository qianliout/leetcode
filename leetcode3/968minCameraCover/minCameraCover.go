package main

import (
	"fmt"
	"math"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Println(minCameraCover(root))

}

func minCameraCover(root *TreeNode) int {
	choose, _, byChild := dfs(root)
	return min(choose, byChild)
}

func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return math.MaxInt32, 0, 0
	}
	lChoose, lByP, lByChild := dfs(root.Left)
	rChoose, rByP, rByChild := dfs(root.Right)

	// 该节点装了，那么两子树装不装都可以
	choose := min(lChoose, lByP, lByChild) + min(rChoose, rByP, rByChild) + 1
	// 该节点没有装，但是被他的父节点监控到了
	byP := min(lChoose, lByChild) + min(rChoose, rByChild)
	// 该节点没有装，被两个儿子监控到了
	byChild := min(lChoose+rByChild, rChoose+lByChild, lChoose+rChoose)
	return choose, byP, byChild
}
