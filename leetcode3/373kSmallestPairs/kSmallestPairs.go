package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 10))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	mh := make(MaxHeap, 0)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if len(mh) < k {
				heap.Push(&mh, Item{N1: n1, N2: n2})
				continue
			}
			top := mh[0]
			if n1+n2 >= top.N1+top.N2 {
				continue
			}
			heap.Pop(&mh)
			heap.Push(&mh, Item{N1: n1, N2: n2})
		}
	}
	ans := make([][]int, 0)
	for len(mh) > 0 {
		pop := heap.Pop(&mh).(Item)
		ans = append(ans, []int{pop.N1, pop.N2})
	}
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}

type Item struct {
	N1 int
	N2 int
}

type MaxHeap []Item

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].N1+h[i].N2 > h[j].N1+h[j].N2 }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
