package main

import (
	"container/heap"
	"fmt"

	. "outback/geeke/leetcode/common/commonHeap"
)

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

// 大顶堆
func maxSlidingWindow1(nums []int, k int) []int {
	ans := make([]int, 0)
	if len(nums) < k {
		return ans
	}
	mheap := make(MaxHeap, 0)
	win := make(map[int]int)
	for i, ch := range nums {
		win[ch]++
		if i >= k {
			win[nums[i-k]]--
		}

		heap.Push(&mheap, ch)
		if i >= k-1 {
			// 取值了
			// 延迟删除
			for {
				top := mheap[0]
				if win[top] > 0 {
					break
				}
				heap.Pop(&mheap)
			}
			ans = append(ans, mheap[0])
		}
	}

	return ans
}

// 优先对列
func maxSlidingWindow2(nums []int, k int) []int {
	ans := make([]int, 0)
	if len(nums) < k || k <= 0 {
		return ans
	}
	que := make(PriorityQueue, 0)
	for i, ch := range nums {
		heap.Push(&que, &IntItem{
			Value:    i,
			Priority: ch,
		})
		if i >= k-1 {
			// 延迟删除
			for {
				top := que[0]
				if top.Value > i-k {
					break
				}
				heap.Pop(&que)
			}
			ans = append(ans, que[0].Priority)
		}
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	if len(nums) < k || k <= 0 {
		return ans
	}
	// 单调递减stark
	stark := make([]int, 0)
	for i, ch := range nums {
		// 移除
		for len(stark) > 0 && i-stark[0] >= k {
			stark = stark[1:]
		}

		for len(stark) > 0 && nums[stark[len(stark)-1]] < ch {
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, i)

		if i >= k-1 {
			ans = append(ans, nums[stark[0]])
		}
	}
	return ans
}
