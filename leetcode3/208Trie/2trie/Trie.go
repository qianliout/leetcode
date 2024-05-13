package main

func main() {

}

type Node struct {
	End  bool
	Next map[byte]*Node
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: &Node{End: false, Next: nil}}
}

func (this *Trie) Insert(word string) {
	insert(this.Root, []byte(word), 0)
}

func (this *Trie) Search(word string) bool {
	return search(this.Root, []byte(word), 0)
}

func (this *Trie) StartsWith(prefix string) bool {

	return startsWith(this.Root, []byte(prefix), 0)
}

func insert(node *Node, word []byte, start int) {
	if start >= len(word) {
		node.End = true
		return
	}
	if node.Next == nil {
		node.Next = make(map[byte]*Node)
	}

	c := word[start]
	if node.Next[c] == nil {
		node.Next[c] = &Node{}
	}
	insert(node.Next[c], word, start+1)
}

func search(node *Node, word []byte, start int) bool {
	if node == nil {
		return false
	}
	if start >= len(word) {
		return node.End
	}
	c := word[start]
	return search(node.Next[c], word, start+1)
}

func startsWith(node *Node, word []byte, start int) bool {
	if node == nil {
		return false
	}
	if start >= len(word) {
		return true
	}
	c := word[start]
	return startsWith(node.Next[c], word, start+1)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
