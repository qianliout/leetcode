package main

import (
	"fmt"
)

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}

// 会超时
func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return make([]string, 0)
	}
	trie := &Tire{Dump: &Node{}}
	for _, ch := range words {
		trie.Insert(ch)
	}
	path := make([]byte, 0)
	ans := make([]string, 0)
	res := make(map[string]bool)
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	mem := make(map[string]bool)
	for i := range board {
		for j := range board[i] {
			dfs(board, i, j, path, res, trie, used, mem)
		}
	}
	for key := range res {
		ans = append(ans, key)
	}

	return ans
}

func dfs(board [][]byte, col, row int, path []byte, ans map[string]bool, trie *Tire, used [][]bool, mem map[string]bool) {
	if col < 0 || col >= len(board) || row < 0 || row >= len(board[col]) {
		return
	}
	if used[col][row] {
		return
	}
	used[col][row] = true
	by := board[col][row]
	path = append(path, by)
	fin, ok := mem[string(path)]
	if ok && fin {
		ans[string(path)] = true
	}
	if !ok {
		if trie.Search(string(path)) {
			ans[string(path)] = true
		}
	}

	dfs(board, col+1, row, path, ans, trie, used, mem)
	dfs(board, col-1, row, path, ans, trie, used, mem)
	dfs(board, col, row+1, path, ans, trie, used, mem)
	dfs(board, col, row-1, path, ans, trie, used, mem)
	used[col][row] = false
	path = path[:len(path)-1]
}

type Node struct {
	End  bool
	Data map[byte]*Node
}

type Tire struct {
	Dump *Node
}

func (vi *Tire) Insert(word string) {
	add(vi.Dump, []byte(word), 0)
}

func (vi *Tire) Search(word string) bool {
	return search(vi.Dump, []byte(word), 0)
}

func (vi *Tire) StartWith(word string) bool {
	return start(vi.Dump, []byte(word), 0)
}

func add(root *Node, word []byte, idx int) {
	if root == nil {
		return
	}
	if idx >= len(word) {
		root.End = true
		return
	}
	by := word[idx]
	if root.Data == nil {
		root.Data = make(map[byte]*Node)
	}
	if root.Data[by] == nil {
		root.Data[by] = &Node{}
	}
	add(root.Data[by], word, idx+1)
}

func search(root *Node, word []byte, idx int) bool {
	if root == nil {
		return false
	}
	if idx >= len(word) {
		return root.End
	}
	by := word[idx]
	if root.Data[by] == nil {
		return false
	}
	return search(root.Data[by], word, idx+1)
}

func start(root *Node, word []byte, idx int) bool {
	if root == nil {
		return false
	}
	if idx >= len(word) {
		return true
	}
	by := word[idx]
	if root.Data[by] == nil {
		return false
	}
	return search(root.Data[by], word, idx+1)
}
