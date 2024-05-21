package main

import (
	"container/list"
	"math"
)

func main() {

}

type Node struct {
	key   int
	value int
	freq  int
}

func NewNode(key, value int, freq int) *Node {
	return &Node{
		key:   key,
		value: value,
		freq:  freq,
	}
}

type LFUCache struct {
	Data    map[int]*list.Element
	Freq    map[int]*list.List
	Cap     int
	MinFreq int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		Data: make(map[int]*list.Element),
		Freq: make(map[int]*list.List),
		Cap:  capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.Data[key]
	if !ok {
		return -1
	}
	no := el.Value.(*Node)
	this.Freq[no.freq].Remove(el)
	// 更新 minFreq
	// 假如MinFreq=1，这个链表里只有一个元素，那么更新后MinFreq 就等于2
	// 会不会存在 Freq=2的链表为空呢？不会，因为 no会马上更新
	if this.MinFreq == no.freq && this.Freq[no.freq].Len() == 0 {
		this.MinFreq = no.freq + 1
	}
	no.freq++
	if this.Freq[no.freq] == nil {
		this.Freq[no.freq] = list.New()
	}
	front := this.Freq[no.freq].PushFront(no)
	this.Data[key] = front

	return no.value
	// 找最小的
}

func (this *LFUCache) Put(key int, value int) {
	if el, ok := this.Data[key]; ok {
		no := el.Value.(*Node)
		this.Freq[no.freq].Remove(el)
		if this.MinFreq == no.freq && this.Freq[no.freq].Len() == 0 {
			this.MinFreq = no.freq + 1
		}

		no.freq++
		no.value = value
		if this.Freq[no.freq] == nil {
			this.Freq[no.freq] = list.New()
		}
		front := this.Freq[no.freq].PushFront(no)
		this.Data[key] = front
		return
	}

	if len(this.Data) >= this.Cap {
		back := this.Freq[this.MinFreq].Back()
		backN := back.Value.(*Node)
		this.Freq[this.MinFreq].Remove(back)
		delete(this.Data, backN.key)
	}
	this.MinFreq = 1
	no := NewNode(key, value, 1)
	if this.Freq[no.freq] == nil {
		this.Freq[no.freq] = list.New()
	}
	front := this.Freq[no.freq].PushFront(no)
	this.Data[key] = front
	return
}

func (this *LFUCache) GetMinFreq() int {
	mi := math.MaxInt64
	for i := range this.Freq {
		li := this.Freq[i]
		if li.Len() > 0 && i < mi {
			mi = i
		}
	}
	if mi == math.MaxInt64 {
		mi = 0
	}
	return mi
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
