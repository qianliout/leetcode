package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	exit := make(map[int]int) // key-->index
	for i, ch := range nums {
		pre, ok := exit[ch]
		if !ok {
			exit[ch] = i
		} else {
			if i-pre <= k {
				return true
			}
			exit[ch] = i
		}
	}

	return false
}
