package main

import (
	"container/list"
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

type Entry struct {
	Key   int
	Value int
}

type LRUCache struct {
	Used int
	Cap  int
	Head *list.List            // must add a dump node
	Data map[int]*list.Element // key-->node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cap:  capacity,
		Data: make(map[int]*list.Element),
		Head: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	entry := this.GetEntry(key)
	if entry != nil {
		return entry.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	en := &Entry{
		Key:   key,
		Value: value,
	}
	entry := this.GetEntry(key)
	if entry != nil {
		entry.Value = value
		return
	}

	if len(this.Data) >= this.Cap {
		back := this.Head.Back()
		this.Head.Remove(back)
		delete(this.Data, back.Value.(*Entry).Key)
	}
	this.Data[key] = this.Head.PushFront(en)
}

func (this *LRUCache) GetEntry(key int) *Entry {
	element, ok := this.Data[key]
	if !ok || element == nil {
		return nil
	}
	en := element.Value.(*Entry)

	this.Head.Remove(element)
	front := this.Head.PushFront(en)
	this.Data[key] = front

	return front.Value.(*Entry)
}

func (this *LRUCache) PutEntry(e *Entry) {
	element := this.Data[e.Key]
	if element != nil {
		this.Head.Remove(element)
		this.Data[e.Key] = this.Head.PushFront(e)

		return
	}
	this.Data[e.Key] = this.Head.PushFront(e)
}
