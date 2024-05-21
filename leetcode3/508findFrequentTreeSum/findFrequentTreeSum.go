package main

import (
	. "outback/geeke/leetcode/common/treenode"
)

func main() {

}

func findFrequentTreeSum(root *TreeNode) []int {
	mem := make(map[*TreeNode]int)
	ans := make(map[int]int)

	dfs(root, mem, ans)

	ma := 0
	for _, v := range ans {
		if v > ma {
			ma = v
		}
	}

	res := make([]int, 0)
	for k, v := range ans {
		if v == ma {
			res = append(res, k)
		}
	}
	return res
}

func dfs(root *TreeNode, mem map[*TreeNode]int, ans map[int]int) int {
	if root == nil {
		return 0
	}
	if va, ok := mem[root]; ok {
		ans[va]++
		return va
	}
	left := dfs(root.Left, mem, ans)
	right := dfs(root.Right, mem, ans)
	mem[root] = root.Val + left + right
	ans[mem[root]]++
	return mem[root]
}
