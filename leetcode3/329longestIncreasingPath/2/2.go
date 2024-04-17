package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	fmt.Println(longestIncreasingPath(matrix))
}

func longestIncreasingPath(matrix [][]int) int {
	ans := 0

	for i := range matrix {
		for j := range matrix[i] {
			used := make(map[[2]int]bool)
			dfs(matrix, i, j, 0, 0, used, &ans)
		}
	}
	return ans
}

// 这种做法是错的，原因是啥呢
func dfs(matrix [][]int, col, row, pre, path int, used map[[2]int]bool, ans *int) {
	if col < 0 || row < 0 || col >= len(matrix) || row >= len(matrix[col]) {
		return
	}
	if matrix[col][row] <= pre {
		if *ans < path-1 {
			fmt.Println("path,", path)
			*ans = path - 1
		}
		return
	}

	a := [2]int{col, row}
	for i := col; i < len(matrix); i++ {
		for j := row; j < len(matrix[i]); j++ {
			if used[a] {
				continue
			}
			// 在下一步才判断
			dfs(matrix, i+1, j, matrix[i][j], path+1, used, ans)
			dfs(matrix, i-1, j, matrix[i][j], path+1, used, ans)
			dfs(matrix, i, j+1, matrix[i][j], path+1, used, ans)
			dfs(matrix, i, j-1, matrix[i][j], path+1, used, ans)
			used[a] = false
		}
	}
}
