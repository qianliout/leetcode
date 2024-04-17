package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	matrix := [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}
	fmt.Println(longestIncreasingPath(matrix))
}

func longestIncreasingPath(matrix [][]int) int {
	ans := 0
	mem := make(map[[2]int]int)
	for i := range matrix {
		for j := range matrix[i] {
			ans = Max(ans, dfs2(matrix, i, j, mem))
		}
	}
	return ans
}

// 可以实现，但是有太多的重复计算
func dfs(matrix [][]int, col, row, path int, used map[[2]int]bool, ans *int) {
	dir := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < 4; i++ {
		newC := col + dir[i][0]
		newR := row + dir[i][1]
		*ans = Max(*ans, path)

		if in(matrix, newC, newR) && matrix[newC][newR] > matrix[col][row] {
			dfs(matrix, newC, newR, path+1, used, ans)
		}
	}
}

// 表示以 col,row 为启点，延伸时的最大值
func dfs2(matrix [][]int, col, row int, mem map[[2]int]int) int {
	if va, ok := mem[[2]int{col, row}]; ok {
		return va
	}

	ans := 1
	dir := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i := 0; i < 4; i++ {
		newC := col + dir[i][0]
		newR := row + dir[i][1]
		if in(matrix, newC, newR) && matrix[newC][newR] > matrix[col][row] {
			ans = Max(ans, dfs2(matrix, newC, newR, mem)+1)
		}
	}
	mem[[2]int{col, row}] = ans
	return ans
}

func in(matrix [][]int, col, row int) bool {
	if col < 0 || row < 0 || col >= len(matrix) || row >= len(matrix[0]) {
		return false
	}
	return true
}
