package main

import (
	"fmt"
)

func main() {
	fmt.Println(orangesRotting([][]int{{0, 2}}))
	fmt.Println(orangesRotting([][]int{{0}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

/*
值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
*/

type Grid struct {
	Col int
	Row int
}

func orangesRotting(grid [][]int) int {
	queue := make([]Grid, 0)
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == 2 {
				queue = append(queue, Grid{Col: i, Row: j})
			}
		}
	}
	ans := 0

	for len(queue) > 0 {
		lev := make([]Grid, 0)
		for i := 0; i < len(queue); i++ {
			first := queue[i]
			res := find(grid, first.Col, first.Row)
			lev = append(lev, res...)
		}
		if len(lev) > 0 {
			ans++
		}
		queue = lev
	}
	for i := range grid {
		for _, ch := range grid[i] {
			if ch == 1 {
				return -1
			}
		}
	}

	return ans
}

// 找烂橘子周边的好橘子
func find(grid [][]int, col, row int) []Grid {
	res := make([]Grid, 0)
	if !in(grid, col, row) {
		return res
	}
	if grid[col][row] != 2 {
		return res
	}
	if in(grid, col-1, row) && grid[col-1][row] == 1 {
		grid[col-1][row] = 2
		res = append(res, Grid{Col: col - 1, Row: row})
	}
	if in(grid, col+1, row) && grid[col+1][row] == 1 {
		grid[col+1][row] = 2
		res = append(res, Grid{Col: col + 1, Row: row})
	}
	if in(grid, col, row+1) && grid[col][row+1] == 1 {
		grid[col][row+1] = 2
		res = append(res, Grid{Col: col, Row: row + 1})
	}
	if in(grid, col, row-1) && grid[col][row-1] == 1 {
		grid[col][row-1] = 2
		res = append(res, Grid{Col: col, Row: row - 1})
	}
	return res
}

func in(grid [][]int, col, row int) bool {
	if col < 0 || col >= len(grid) || row < 0 || row >= len(grid[col]) {
		return false
	}
	return true
}
