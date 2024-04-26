package main

import (
	"container/list"
)

func main() {
	all := Constructor()
	all.Inc("hello")
	all.Inc("hello")
}

type Entry struct {
	key  string
	freq int
}

type AllOne struct {
	keyToNode  map[string]*list.Element
	freqToList map[int]*list.List // key-->频率
}

func Constructor() AllOne {
	return AllOne{
		keyToNode:  make(map[string]*list.Element),
		freqToList: make(map[int]*list.List),
	}
}

func (this *AllOne) Inc(key string) {

	node, ok := this.keyToNode[key]
	if ok {
		en := &Entry{
			key:  key,
			freq: 0,
		}

	}

	this.Data[key]++
	if this.Data[key] > this.MinCnt {
		this.Max = key
		this.MaxCnt = this.Data[key]
	}
	if this.Min == "" {
		this.Max = key
		this.MaxCnt = this.Data[key]
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.Data[key]; !ok {
		return
	}
	this.Data[key]--
	if this.Data[key] < this.MinCnt {
		this.Min = key
		this.MinCnt = this.Data[key]
	}
}

func (this *AllOne) GetMaxKey() string {
	return this.Max

}

func (this *AllOne) GetMinKey() string {
	return this.Min
}
