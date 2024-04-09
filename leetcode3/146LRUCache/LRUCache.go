package main

import (
	"fmt"
)

func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(3))
}

type Node struct {
	Key   int
	Value int
	Next  *Node
	Pre   *Node
}

type LRUCache struct {
	Used int
	Cap  int
	Head *Node         // must add a dump node
	Data map[int]*Node // key-->node
}

func Constructor(capacity int) LRUCache {
	dump := &Node{}
	return LRUCache{
		Cap:  capacity,
		Data: make(map[int]*Node),
		Head: dump,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Data[key]; ok {
		this.Head.RemoveNode(node)
		this.Head.AddFirst(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Data[key]; ok {
		node.Value = value
		this.Head.RemoveNode(node)
		this.Head.AddFirst(node)
		return
	}
	node := &Node{Key: key, Value: value}

	if this.Used < this.Cap {
		this.Head.AddFirst(node)
		this.Used++
	} else {
		pop := this.Head.PopLast()
		if pop != nil {
			delete(this.Data, pop.Key)
		}
		this.Head.AddFirst(node)
	}
	this.Data[key] = node
}

func (vi *Node) AddFirst(node *Node) {
	nx := vi.Next
	if nx != nil {
		nx.Pre = node
	}

	node.Pre = vi
	node.Next = nx
	vi.Next = node
}

func (vi *Node) RemoveNode(node *Node) {
	if node == nil {
		return
	}
	pre := node.Pre
	nx := node.Next
	if nx != nil {
		nx.Pre = pre
	}

	pre.Next = nx
}

func (vi *Node) PopLast() *Node {
	cur := vi
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}

	if cur != nil {
		cur.Pre.Next = nil
		return cur
	}
	return nil
}
