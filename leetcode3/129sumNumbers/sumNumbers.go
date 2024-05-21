package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(sumNumbers(root))
}

func sumNumbers2(root *TreeNode) int {
	pre := 0
	res := 0
	dfs(root, pre, &res)
	return res
}

// will change every node's value
func dfs(root *TreeNode, pre int, res *int) {
	if root == nil {
		return
	}
	root.Val = root.Val + pre*10
	if root.Left == nil && root.Right == nil {
		*res = *res + root.Val
	}
	dfs(root.Left, root.Val, res)
	dfs(root.Right, root.Val, res)
}

func sumNumbers(root *TreeNode) int {

	path := make([]int, 0)
	// used := make(map[*TreeNode]bool)
	ans := 0
	backtrace2(root, path, &ans)
	return ans
}

// not change node value
func backtrace(root *TreeNode, path []int, used map[*TreeNode]bool, ans *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*ans += sumPath(path)
		path = path[:len(path)-1]
		return
	}
	if used[root] {
		return
	}
	used[root] = true
	path = append(path, root.Val)
	backtrace(root.Left, path, used, ans)
	backtrace(root.Right, path, used, ans)
	path = path[:len(path)-1]
	used[root] = false
}

func backtrace2(root *TreeNode, path []int, ans *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val)
		*ans += sumPath(path)
		path = path[:len(path)-1]
		return
	}
	path = append(path, root.Val)
	backtrace2(root.Left, path, ans)
	backtrace2(root.Right, path, ans)
	path = path[:len(path)-1]
}

func sumPath(nums []int) int {
	ans := 0
	for i := range nums {
		ans = ans*10 + nums[i]
	}
	return ans
}
