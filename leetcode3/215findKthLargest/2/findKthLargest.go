package main

import (
	"container/heap"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {

}

func findKthLargest(nums []int, k int) int {
	if k <= 0 {
		return 0
	}
	he := make(MinHeap, 0)
	for _, ch := range nums {
		if len(he) < k {
			heap.Push(&he, ch)
			continue
		}

		top := he[0]
		if top >= ch {
			continue
		}
		heap.Pop(&he)
		heap.Push(&he, ch)
	}
	return he[0]
}
