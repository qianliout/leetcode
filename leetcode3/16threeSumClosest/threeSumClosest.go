package main

import (
	"math"
	"sort"

	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := math.MaxInt64
	sub := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tag := nums[left] + nums[right] + nums[i]
			if tag == target {
				return tag
			} else if tag > target {
				right--
			} else if tag < target {
				left++
			}

			if AbsSub(tag, target) < sub {
				sub = AbsSub(tag, target)
				res = tag
			}
		}
	}

	return res
}
