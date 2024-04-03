package main

import (
	"fmt"
)

func main() {
	// board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	board := [][]byte{[]byte("OOO"), []byte("OOO"), []byte("OOO")}

	solve(board)

	for i := range board {
		fmt.Println(string(board[i]))
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board)-1, len(board[0])-1
	for col := 0; col <= m; col++ {
		if board[col][0] == 'O' {
			full(board, col, 0)
		}
		if board[col][n] == 'O' {
			full(board, col, n)
		}
	}

	for row := 0; row <= n; row++ {
		if board[0][row] == 'O' {
			full(board, 0, row)
		}
		if board[m][row] == 'O' {
			full(board, m, row)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

func full(grid [][]byte, col, row int) {
	if col < 0 || col >= len(grid) || row < 0 || row >= len(grid[0]) {
		return
	}

	if grid[col][row] == 'O' {
		grid[col][row] = 'Y'
		full(grid, col+1, row)
		full(grid, col-1, row)
		full(grid, col, row+1)
		full(grid, col, row-1)
	}
}
