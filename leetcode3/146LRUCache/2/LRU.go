package main

import (
	"container/list"
	"fmt"
)

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
	lru.Put(4, 4)
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(4))
}

type Entry struct {
	Key   int
	Value int
}

type LRUCache struct {
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

func (vi *LRUCache) Get(key int) int {
	el := vi.Data[key]
	if el == nil {
		return -1
	}
	en := el.Value.(*Entry)
	vi.Head.Remove(el)
	vi.Data[en.Key] = vi.Head.PushFront(en)
	// 这样写也可以
	// vi.Head.MoveToFront(el)

	return en.Value
}

func (vi *LRUCache) Put(key int, value int) {
	en := &Entry{
		Key:   key,
		Value: value,
	}

	el := vi.Data[key]
	// 已存在
	if el != nil {
		vi.Head.Remove(el)
		front := vi.Head.PushFront(en)
		vi.Data[key] = front
		return
	}

	if len(vi.Data) >= vi.Cap {

		back := vi.Head.Back()
		last := back.Value.(*Entry)

		delete(vi.Data, last.Key)

		vi.Head.Remove(back)

		front := vi.Head.PushFront(en)
		vi.Data[key] = front
		return
	}

	if len(vi.Data) < vi.Cap {
		front := vi.Head.PushFront(en)
		vi.Data[key] = front
		return
	}
}
