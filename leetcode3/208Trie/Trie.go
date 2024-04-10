package main

import (
	"fmt"
)

func main() {
	trie := Constructor()
	trie.Insert("hello")
	trie.Insert("word")
	trie.Insert("fuck")
	fmt.Println(trie.Search("word"))
	fmt.Println(trie.Search("wor"))
	fmt.Println(trie.StartsWith("wor"))
}

type Node struct {
	End bool
	// Key   byte
	Value map[byte]*Node
}

type Trie struct {
	Dump *Node
}

func Constructor() Trie {

	return Trie{
		Dump: &Node{},
	}
}

func (this *Trie) Insert(word string) {
	add(this.Dump, []byte(word), 0)
}

func (this *Trie) Search(word string) bool {
	return search(this.Dump, []byte(word), 0)
}

func (this *Trie) StartsWith(prefix string) bool {
	return searchPrefix(this.Dump, []byte(prefix), 0)
}

func add(root *Node, word []byte, idx int) {
	if idx >= len(word) {
		return
	}
	key := word[idx]
	if root.Value == nil {
		root.Value = make(map[byte]*Node)
	}

	if root.Value[key] == nil {
		root.Value[key] = &Node{}
	}
	node := root.Value[key]

	if idx == len(word)-1 {
		node.End = true
		return
	}
	add(node, word, idx+1)
}

func search(dump *Node, word []byte, idx int) bool {

	if idx >= len(word) || dump == nil {
		return false
	}
	key := word[idx]
	node := dump.Value[key]
	if node == nil {
		return false
	}
	if idx == len(word)-1 {
		return node.End
	}
	return search(node, word, idx+1)
}

func searchPrefix(dump *Node, word []byte, idx int) bool {
	if idx >= len(word) || dump == nil {
		return false
	}
	key := word[idx]
	node := dump.Value[key]
	if node == nil {
		return false
	}
	if idx == len(word)-1 {
		return true
	}

	return searchPrefix(node, word, idx+1)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
