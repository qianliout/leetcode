package main

import (
	"fmt"
)

func main() {
	trie := Constructor()
	trie.AddWord("hello")
	trie.AddWord("word")
	trie.AddWord("fuck")
	fmt.Println(trie.Search("word"))
	fmt.Println(trie.Search("wor"))
	fmt.Println(trie.Search("wor."))
}

type WordDictionary struct {
	Dump *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Dump: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	add(this.Dump, []byte(word), 0)
}

func (this *WordDictionary) Search(word string) bool {
	return search(this.Dump, []byte(word), 0)
}

type Node struct {
	End  bool
	Data map[byte]*Node
}

type Tire struct {
	Dump *Node
}

func add(root *Node, word []byte, idx int) {
	if root == nil {
		return
	}
	if idx >= len(word) {
		root.End = true
		return
	}
	if root.Data == nil {
		root.Data = make(map[byte]*Node)
	}
	by := word[idx]
	if root.Data[by] == nil {
		root.Data[by] = &Node{}
	}
	no := root.Data[by]
	add(no, word, idx+1)
}

func search(root *Node, word []byte, idx int) bool {
	if root == nil {
		return false
	}
	if idx >= len(word) {
		return root.End

	}
	by := word[idx]

	if by != '.' {
		if root.Data[by] == nil {
			return false
		}
		no := root.Data[by]
		return search(no, word, idx+1)
	} else {
		var exit bool
		for _, node := range root.Data {
			if search(node, word, idx+1) {
				exit = true
			}
		}
		return exit
	}
}
