package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMoves2([]int{1, 0, 0, 8, 6}))
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	mid := nums[len(nums)/2]

	cnt := 0
	for _, ch := range nums {
		if ch > mid {
			cnt += ch - mid
		} else {
			cnt += mid - ch
		}
	}

	return cnt
}
