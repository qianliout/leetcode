package main

import (
	"container/heap"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 || k <= 0 {
		return 0
	}

	mh := make(MaxHeap, 0)

	for i := range matrix {
		for _, ch := range matrix[i] {
			if len(mh) < k {
				heap.Push(&mh, ch)
				continue
			}
			top := mh[0]
			if ch < top {
				heap.Pop(&mh)
				heap.Push(&mh, ch)
			}
		}
	}
	return mh[0]
}
