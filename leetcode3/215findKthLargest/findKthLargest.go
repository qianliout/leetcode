package main

import (
	"container/heap"
	"fmt"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k <= 0 {
		return 0
	}

	// 小顶堆

	he := make(MinHeap, 0)
	for _, ch := range nums {
		heap.Push(&he, ch)

		for len(he) > k {
			heap.Pop(&he)
		}
	}

	return he[0]
}
