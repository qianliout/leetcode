package main

import (
	"container/heap"
	"fmt"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {
	md := Constructor()
	md.AddNum(6)
	md.AddNum(7)
	md.AddNum(8)
	fmt.Println(md.FindMedian())

	md.AddNum(1)
	fmt.Println(md.FindMedian())
	md.AddNum(2)

	fmt.Println(md.FindMedian())
}

type MedianFinder struct {
	Right MinHeap
	Left  MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		Right: make(MinHeap, 0),
		Left:  make(MaxHeap, 0),
	}
}
func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.Left, num)
	// 调整

	for len(this.Left)-len(this.Right) > 1 || (len(this.Left) > 0 && len(this.Right) > 0 && this.Left[0] > this.Right[0]) {
		if len(this.Left)-len(this.Right) > 1 {
			heap.Push(&this.Right, heap.Pop(&this.Left).(int))
		}
		if len(this.Left) > 0 && len(this.Right) > 0 && this.Left[0] > this.Right[0] {
			heap.Push(&this.Left, heap.Pop(&this.Right).(int))
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Left) == 0 {
		return 0
	}
	if len(this.Left) > len(this.Right) {
		return float64(this.Left[0])
	}

	return float64(this.Left[0]+this.Right[0]) / 2
}
