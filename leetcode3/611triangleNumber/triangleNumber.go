package main

import (
	"sort"
)

func main() {

}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if check(nums[i], nums[j], nums[k]) {
					ans++
				} else {
					break
				}
			}
		}
	}
	return ans
}

func check(a, b, c int) bool {
	if a+b <= c {
		return false
	}
	return true
}
