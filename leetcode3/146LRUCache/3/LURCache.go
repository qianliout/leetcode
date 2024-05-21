package main

import (
	"container/list"
)

func main() {

}

type Node struct {
	key   int
	value int
}

type LRUCache struct {
	Data map[int]*list.Element
	Head *list.List
	Cap  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Data: make(map[int]*list.Element),
		Head: list.New(),
		Cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.Data[key]
	if !ok {
		return -1
	}
	no := el.Value.(*Node)
	this.Head.MoveToFront(el)
	return no.value
}

func (this *LRUCache) Put(key int, value int) {
	el, ok := this.Data[key]
	if ok {
		no := &Node{key: key, value: value}
		this.Head.Remove(el)
		front := this.Head.PushFront(no)
		this.Data[key] = front
		return
	}
	if this.Head.Len() < this.Cap {
		no := &Node{key: key, value: value}
		front := this.Head.PushFront(no)
		this.Data[key] = front
		return
	}
	back := this.Head.Back()
	backN := back.Value.(*Node)

	delete(this.Data, backN.key)

	this.Head.Remove(back)
	no := &Node{key: key, value: value}
	front := this.Head.PushFront(no)
	this.Data[key] = front
	return
}
