package main

import (
	"fmt"

	. "outback/geeke/leetcode/common/utils"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}

func findClosestElements(arr []int, k int, x int) []int {
	if k >= len(arr) {
		return arr
	}
	left, right := 0, len(arr)-1

	redu := len(arr) - k

	for redu > 0 {
		if AbsSub(arr[left], x) > AbsSub(arr[right], x) {
			left++
		} else {
			right--
		}
		redu--
	}
	return arr[left : right+1]
}
