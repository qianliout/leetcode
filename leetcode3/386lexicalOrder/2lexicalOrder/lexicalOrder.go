package main

import (
	"fmt"
)

func main() {
	fmt.Println(lexicalOrder(29))
}

func lexicalOrder(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= 9; i++ {
		if len(res) >= n {
			break
		}
		res = append(res, i)
		dfs(i, n, &res)
	}
	return res
}

func dfs(pre, limit int, res *[]int) {
	if pre > limit {
		return
	}
	for i := 0; i <= 9; i++ {
		nu := pre*10 + i
		if nu > limit {
			break
		}
		*res = append(*res, nu)
		dfs(nu, limit, res)
	}
}
