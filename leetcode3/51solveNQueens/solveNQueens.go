package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	qu := make([][]byte, n)
	for i := range qu {
		qu[i] = make([]byte, n)
		for j := range qu[i] {
			qu[i][j] = '.'
		}
	}
	ans := make([][]string, 0)
	dfs(qu, 0, n, &ans)
	return ans
}

func dfs(queens [][]byte, col, n int, ans *[][]string) {
	if col >= n {
		*ans = append(*ans, gen(queens))
		return
	}
	for i := 0; i < n; i++ {
		if check(queens, col, i, n) {
			queens[col][i] = 'Q'
			dfs(queens, col+1, n, ans)
			queens[col][i] = '.'
		}
	}
}

func check(path [][]byte, col, row, n int) bool {
	// 检查列(向上)
	for i := 0; i < col; i++ {
		if path[i][row] == 'Q' {
			return false
		}
	}
	// 再检查行(向左)
	for i := 0; i < row; i++ {
		if path[col][i] == 'Q' {
			return false
		}
	}
	// 再检查45度
	for cl, ro := col-1, row+1; cl >= 0 && ro < n; cl, ro = cl-1, ro+1 {
		if path[cl][ro] == 'Q' {
			return false
		}
	}

	// 再检查135度
	for cl, ro := col-1, row-1; cl >= 0 && ro >= 0; cl, ro = cl-1, ro-1 {
		if path[cl][ro] == 'Q' {
			return false
		}
	}
	return true
}

func gen(queens [][]byte) []string {
	ans := make([]string, 0)
	for i := range queens {
		ans = append(ans, string(queens[i]))
	}

	return ans
}
