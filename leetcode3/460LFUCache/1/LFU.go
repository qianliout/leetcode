package main

import (
	"container/list"
)

func main() {
	lrf := Constructor(2)
	lrf.Put(1, 1)
	lrf.Put(2, 2)
	lrf.Get(1)
	lrf.Put(3, 3)
}

type Entry struct {
	key   int
	value int
	freq  int
}

type LFUCache struct {
	capacity   int
	minFreq    int
	freqToList map[int]*list.List
	key2Node   map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		key2Node:   map[int]*list.Element{},
		freqToList: map[int]*list.List{},
	}
}

func (vi *LFUCache) Put(key, value int) {
	en := &Entry{
		key:   key,
		value: value,
		freq:  1,
	}

	// 如果 key 存在
	if el, ok := vi.key2Node[key]; ok && el != nil {
		preEn := el.Value.(*Entry)

		li1 := vi.freqToList[preEn.freq]

		li1.Remove(el)

		if li1.Len() == 0 && vi.minFreq == preEn.freq {
			delete(vi.freqToList, preEn.freq)
			vi.minFreq = preEn.freq + 1
		}

		// 如果原来就已经存在这个 key,那么这个 key 的 freq 要加1
		en.freq = preEn.freq + 1

		if vi.freqToList[en.freq] == nil {
			vi.freqToList[en.freq] = list.New()
		}
		el2 := vi.freqToList[en.freq].PushFront(en)
		vi.key2Node[key] = el2
		return
	}

	// 如果 key 不存在
	// 1 还有空间
	if vi.capacity > len(vi.key2Node) {
		vi.minFreq = en.freq

		if vi.freqToList[en.freq] == nil {
			vi.freqToList[en.freq] = list.New()
		}

		el := vi.freqToList[en.freq].PushFront(en)
		vi.key2Node[key] = el
		return
	}

	// 没有空间了
	if vi.capacity <= len(vi.key2Node) {

		// 先删除一个
		li := vi.freqToList[vi.minFreq]

		el2 := li.Back()
		en2 := el2.Value.(*Entry)

		delete(vi.key2Node, en2.key)

		li.Remove(el2)

		if li.Len() == 0 {
			delete(vi.freqToList, vi.minFreq)
		}

		if vi.freqToList[en.freq] == nil {
			vi.freqToList[en.freq] = list.New()
		}

		front := vi.freqToList[en.freq].PushFront(en)
		vi.key2Node[en.key] = front
		vi.minFreq = en.freq
		return
	}
}

func (vi *LFUCache) Get(key int) int {
	el, ok := vi.key2Node[key]
	if !ok || el == nil {
		return -1
	}

	en := el.Value.(*Entry)
	// 把原 en.freq 删除,上面已 经判断了中否存在，这里可以不用判断
	vi.freqToList[en.freq].Remove(el)
	if vi.freqToList[en.freq].Len() == 0 {
		delete(vi.freqToList, en.freq)
		if vi.minFreq == en.freq {
			// 如果这个 key 的 频率刚好是最小
			vi.minFreq++
		}
	}

	// 频率加1之后，再加入链表
	en.freq++
	if vi.freqToList[en.freq] == nil {
		vi.freqToList[en.freq] = list.New()
	}
	el2 := vi.freqToList[en.freq].PushFront(en)
	vi.key2Node[key] = el2
	return en.value
}
