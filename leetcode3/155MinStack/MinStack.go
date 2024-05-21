package main

import (
	"container/heap"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {

}

// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
type MinStack struct {
	MinHeap MinHeap
	Poped   map[int]int
	Stark   []int
}

func Constructor() MinStack {
	return MinStack{
		MinHeap: make(MinHeap, 0),
		Poped:   make(map[int]int),
		Stark:   make([]int, 0),
	}

}

func (this *MinStack) Push(val int) {
	this.Stark = append(this.Stark, val)
	heap.Push(&this.MinHeap, val)
}

func (this *MinStack) Pop() {
	if len(this.Stark) == 0 {
		return
	}
	pop := this.Stark[len(this.Stark)-1]
	this.Poped[pop]++
	this.Stark = this.Stark[:len(this.Stark)-1]
}

func (this *MinStack) Top() int {
	if len(this.Stark) == 0 {
		return 0
	}
	return this.Stark[len(this.Stark)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Stark) == 0 {
		return 0
	}
	for len(this.MinHeap) > 0 {
		mi := this.MinHeap[0]
		if this.Poped[mi] > 0 {
			this.Poped[mi]--
			heap.Pop(&this.MinHeap)
			if this.Poped[mi] == 0 {
				delete(this.Poped, mi)
			}
			continue
		}
		return mi
	}
	return 0
}
