package main

func main() {

}

// 会超时
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	ans := false
	used := make([][]bool, len(board))
	for k := range used {
		used[k] = make([]bool, len(board[0]))
	}

	for i := range board {
		for j := range board[i] {
			dfs(board, i, j, used, []byte{}, word, &ans)
		}
	}

	return ans
}

func dfs(board [][]byte, col, row int, used [][]bool, path []byte, word string, ans *bool) {
	if *ans {
		return
	}
	if string(path) == word {
		*ans = true
		return
	}

	if col < 0 || row < 0 || col >= len(board) || row >= len(board[col]) {
		return
	}

	if used[col][row] {
		return
	}
	// 没有这一步剪枝就会超时
	if len(path) >= len(word) {
		return
	}
	path = append(path, board[col][row])

	used[col][row] = true
	dfs(board, col+1, row, used, path, word, ans)
	dfs(board, col-1, row, used, path, word, ans)
	dfs(board, col, row+1, used, path, word, ans)
	dfs(board, col, row-1, used, path, word, ans)
	path = path[:len(path)-1]
	used[col][row] = false
}
