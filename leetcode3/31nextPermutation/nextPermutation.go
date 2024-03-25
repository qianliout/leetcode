package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{2, 1, 3})
	nextPermutation([]int{2, 3, 1})
	nextPermutation([]int{3, 2, 1})
}

func nextPermutation(nums []int) {
	start, end, sub := -1, -1, math.MaxInt64
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			start = i - 1
			break
		}
	}
	if start == -1 {
		le, ri := 0, len(nums)-1
		for le < ri {
			nums[le], nums[ri] = nums[ri], nums[le]
			le++
			ri--
		}
		fmt.Println(nums)
		return
	}
	end = start

	for j := start + 1; j < len(nums); j++ {
		if nums[j] > nums[start] && nums[j]-nums[start] < sub {
			end = j
			sub = nums[j] - nums[start]
		}
	}
	nums[start], nums[end] = nums[end], nums[start]
	if start+1 < len(nums) {
		sort.Ints(nums[start+1:])
	}
	fmt.Println(nums)
}
