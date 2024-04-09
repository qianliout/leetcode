package main

import (
	"fmt"
)

func main() {
	num := [][]int{{1, 2, 3}}
	f(num)
	fmt.Println(num)
}

func f(num [][]int) {

	for i := range num {
		for j := range num[i] {
			num[i][j] = 1
		}
	}
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				// full(grid, i+1, j)
				// full(grid, i-1, j)
				// full(grid, i, j+1)
				// full(grid, i, j-1)
			}
		}
	}
	return ans
}

func full(grid [][]byte, col, row int) {
	if col<0
	if col < 0 || col >= len(grid) || row < 0 || row >= len(grid[0]) {
		return
	}
	if grid[col][row] == '0' {
		return
	}
	grid[col][row] = '0'

	full(grid, col+1, row)
	full(grid, col-1, row)
	full(grid, col, row+1)
	full(grid, col, row-1)
}
