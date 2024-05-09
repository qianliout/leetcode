package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"0", "1101", "01", "00111", "1", "10010", "0", "0", "00", "1", "11", "0011"}, 63, 36))
}

func findMaxForm2(strs []string, m int, n int) int {
	nodes := make([]Node, len(strs))
	for i := range strs {
		nodes[i] = help(strs[i])
	}
	ans := 0
	used := make([]bool, len(strs))
	dfs(nodes, used, 0, m, n, &ans)
	return ans
}

// 可以得到结果，但是会超时
func dfs(strs []Node, used []bool, path int, m, n int, ans *int) {
	*ans = Max(path, *ans)
	for i, ch := range strs {
		if used[i] {
			continue
		}
		if ch.Zero > n || ch.One > m {
			continue
		}
		path++
		used[i] = true
		dfs(strs, used, path, m-ch.One, n-ch.Zero, ans)
		used[i] = false
		path--
	}
}

func help(str string) Node {
	owe, zero := 0, 0
	for i := range str {
		if str[i] == '0' {
			owe++
		} else {
			zero++
		}
	}
	return Node{
		One:  owe,
		Zero: zero,
	}
}

type Node struct {
	One  int
	Zero int
}

func findMaxForm(strs []string, m int, n int) int {
	nodes := make([]Node, len(strs))
	for i := range strs {
		nodes[i] = help(strs[i])
	}
	ans := 0
	used := make([]bool, len(strs))
	dfs(nodes, used, 0, m, n, &ans)
	return ans
}

// 背包
func usedp(strs []Node, m, n int) int {
	dp := make([][][]int, len(strs))
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	// 初值

	for i := 0; i < len(strs); i++ {
		st := strs[i]
		// 不选
		dp[i][]

	}

}
