package main

import (
	"container/list"
)

func main() {

}

type Entry struct {
	key   int
	value int
	freq  int
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

func (s *LFUCache) GetEntry(key int) *Entry {
	ele, ok := s.keyToNode[key]
	if !ok || ele == nil {
		return nil
	}
	entry := ele.Value.(*Entry)

	lst := s.freqToList[entry.freq]

	lst.Remove(ele)

	if lst.Len() == 0 {
		delete(s.freqToList, entry.freq)

		if s.minFreq == entry.freq {
			s.minFreq++
		}
	}

	entry.freq++
	s.PutEntry(entry)
	return entry
}

func (s *LFUCache) PutEntry(e *Entry) {
	node, ok := s.freqToList[e.freq]
	if !ok || node == nil {
		s.freqToList[e.freq] = list.New()
	}
	el := s.freqToList[e.freq].PushFront(e)
	s.keyToNode[e.key] = el
}

func (s *LFUCache) Get(key int) int {
	entry := s.GetEntry(key)
	if entry != nil {
		return entry.value
	}
	return -1
}

func (s *LFUCache) Put(key int, value int) {
	e := &Entry{
		key:   key,
		value: value,
		freq:  1,
	}

	if entry := s.GetEntry(key); entry != nil {
		entry.value = value
		return
	}
	// not in cache
	if len(s.keyToNode) < s.capacity {
		s.PutEntry(e)
		s.minFreq = 1
		return
	}

	if len(s.keyToNode) >= s.capacity {
		lst := s.freqToList[s.minFreq]
		backN := lst.Back()
		if backN != nil {
			delete(s.keyToNode, backN.Value.(*Entry).key)
			lst.Remove(backN)
			if lst.Len() == 0 {
				delete(s.freqToList, s.minFreq)
			}
		}
		s.PutEntry(e)
		s.minFreq = 1
	}
}
