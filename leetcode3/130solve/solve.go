package main

func main() {

}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board)-1, len(board[0])-1
	for col := range board {
		if board[col][0] == 'O' {
			board[col][0] = 'Y'
			full(board, col, 0)
		}
		if board[col][n] == 'O' {
			board[col][n] = 'Y'
			full(board, col, n)
		}
	}

	for row := range board[0] {
		if board[0][row] == 'O' {
			board[0][row] = 'Y'
			full(board, 0, row)
		}
		if board[m][row] == 'O' {
			board[m][row] = 'Y'
			full(board, m, row)
		}
	}

	for i := range board {
		for j := range board {
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
	if grid[col][row] == 'Y' {
		return
	}

	if grid[col][row] == 'X' {
		return
	}
	grid[col][row] = 'X'
	full(grid, col+1, row)
	full(grid, col-1, row)
	full(grid, col, row+1)
	full(grid, col, row-1)
}
