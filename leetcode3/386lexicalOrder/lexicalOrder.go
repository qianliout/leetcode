package main

import (
	"fmt"
)

func main() {
	fmt.Println(lexicalOrder(25))
}

func lexicalOrder(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= 9; i++ {
		dfs(n, i, &res)
	}
	return res
}

func dfs(n, cur int, ans *[]int) {
	if cur > n {
		return
	}
	*ans = append(*ans, cur)
	for i := 0; i <= 9; i++ {
		dfs(n, cur*10+i, ans)
	}
}
