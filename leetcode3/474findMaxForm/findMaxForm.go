package main

import (
	"context"
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"0", "1101", "01", "00111", "1", "10010", "0", "0", "00", "1", "11", "0011"}, 63, 36))

	ctx := context.Background()
	context.WithValue(ctx, "", "")
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

// timeout
func findMaxForm3(strs []string, m int, n int) int {
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
// m 个 0 和 n 个 1
func findMaxForm(strs []string, m int, n int) int {
	nums1 := make([]int, len(strs))
	nums0 := make([]int, len(strs))
	for i := range strs {
		nums1[i] = count1(strs[i])
		nums0[i] = len(strs[i]) - count1(strs[i])
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(strs); i++ {
		for j := m; j >= nums0[i]; j-- {
			for k := n; k >= nums1[i]; k-- {
				dp[j][k] = Max(dp[j][k], dp[j-nums0[i]][k-nums1[i]]+1)
			}
		}
	}

	return dp[m][n]
}

func count1(str string) int {
	ans := 0
	for i := range str {
		if str[i]-48 == 1 {
			ans++
		}
	}
	return ans
}
