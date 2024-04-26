package main

import (
	"container/list"
)

type entry struct {
	key, value, freq int // freq 表示这本书被看的次数
}

type LFUCache struct {
	capacity   int
	minFreq    int
	keyToNode  map[int]*list.Element
	freqToList map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		keyToNode:  map[int]*list.Element{},
		freqToList: map[int]*list.List{},
	}
}

func (c *LFUCache) pushFront(e *entry) {
	if _, ok := c.freqToList[e.freq]; !ok {
		c.freqToList[e.freq] = list.New() // 双向链表
	}
	c.keyToNode[e.key] = c.freqToList[e.freq].PushFront(e)
}

func (c *LFUCache) getEntry(key int) *entry {
	node := c.keyToNode[key]
	if node == nil { // 没有这本书
		return nil
	}
	e := node.Value.(*entry)
	lst := c.freqToList[e.freq]
	lst.Remove(node)    // 把这本书抽出来
	if lst.Len() == 0 { // 抽出来后，这摞书是空的
		delete(c.freqToList, e.freq) // 移除空链表
		if c.minFreq == e.freq {     // 这摞书是最左边的
			c.minFreq++
		}
	}
	e.freq++       // 看书次数 +1
	c.pushFront(e) // 放在右边这摞书的最上面
	return e
}

func (c *LFUCache) Get(key int) int {
	if e := c.getEntry(key); e != nil { // 有这本书
		return e.value
	}
	return -1 // 没有这本书
}

func (c *LFUCache) Put(key, value int) {
	if e := c.getEntry(key); e != nil { // 有这本书
		e.value = value // 更新 value
		return
	}
	if len(c.keyToNode) == c.capacity { // 书太多了
		lst := c.freqToList[c.minFreq]                           // 最左边那摞书
		delete(c.keyToNode, lst.Remove(lst.Back()).(*entry).key) // 移除这摞书的最下面的书
		if lst.Len() == 0 {                                      // 这摞书是空的
			delete(c.freqToList, c.minFreq) // 移除空链表
		}
	}
	c.pushFront(&entry{key, value, 1}) // 新书放在「看过 1 次」的最上面
	c.minFreq = 1
}
