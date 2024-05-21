package main

import (
	"fmt"
)

func main() {

	board := [][]byte{[]byte("ABCE"), []byte("SFCS"), []byte("ADEE")}
	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "ABCB"))

}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	res := make([]string, 0)
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			path := make([]byte, 0)
			dfs(board, path, word, i, j, used, &res)
		}
	}
	return len(res) > 0
}

func dfs(board [][]byte, path []byte, word string, col, row int, used [][]bool, ans *[]string) {
	if col < 0 || col >= len(board) || row < 0 || row >= len(board[0]) {
		return
	}
	if len(*ans) > 0 {
		return
	}

	if !used[col][row] {
		path = append(path, board[col][row])

		if len(path) >= len(word) {
			if string(path) == word {
				*ans = append(*ans, word)
			}
			return
		}

		used[col][row] = true

		dfs(board, path, word, col+1, row, used, ans)
		dfs(board, path, word, col-1, row, used, ans)
		dfs(board, path, word, col, row+1, used, ans)
		dfs(board, path, word, col, row-1, used, ans)
		used[col][row] = false
		path = path[:len(path)-1]
	}
}
