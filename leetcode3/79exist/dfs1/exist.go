package main

func main() {

}

func exist(board [][]byte, word string) bool {
	ans := false

	for i := range board {
		for j := range board[i] {
			used := make([][]bool, len(board))
			for k := range used {
				used[k] = make([]bool, len(board[i]))
			}
			dfs(board, i, j, used, 0, []byte(word), &ans)
		}
	}

	return ans
}

func dfs(board [][]byte, col, row int, used [][]bool, idx int, word []byte, ans *bool) {
	if idx >= len(word) {
		*ans = true
		return
	}
	if col < 0 || row < 0 || col >= len(board) || row >= len(board[col]) {
		return
	}
	if *ans {
		return
	}
	if board[col][row] != word[idx] {
		return
	}

	if used[col][row] {
		return
	}

	used[col][row] = true
	dfs(board, col+1, row, used, idx+1, word, ans)
	dfs(board, col-1, row, used, idx+1, word, ans)
	dfs(board, col, row+1, used, idx+1, word, ans)
	dfs(board, col, row-1, used, idx+1, word, ans)
	used[col][row] = false
}
