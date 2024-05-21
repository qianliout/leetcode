package main

import (
	"container/heap"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {

}

func topKFrequent(nums []int, k int) []int {
	// key->freq,每个数值的频率
	mem := make(map[int]int)

	for _, ch := range nums {
		mem[ch]++
	}
	ph := make(PriorityQueue, 0)

	for k, v := range mem {
		heap.Push(&ph, &IntItem{Value: k, Priority: v})
	}

	ans := make([]int, 0)
	for len(ph) > 0 && k > 0 {
		pop := heap.Pop(&ph).(*IntItem)
		ans = append(ans, pop.Value)
		k--
	}
	return ans
}
