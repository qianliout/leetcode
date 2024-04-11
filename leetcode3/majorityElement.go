package main

import (
	"fmt"
)

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func majorityElement(nums []int) []int {
	ans := make([]int, 0)

	if len(nums) == 0 {
		return ans
	}

	ca1, cn1 := nums[0], 0
	ca2, cn2 := nums[0], 0

	for _, ch := range nums {
		if ch == ca1 {
			cn1++
			continue
		}
		if ch == ca2 {
			cn2++
			continue
		}
		if cn1 == 0 {
			ca1 = ch
			cn1 = 1
			continue
		}
		if cn2 == 0 {
			ca2 = ch
			cn2 = 1
			continue
		}
		cn1--
		cn2--
	}
	cn1, cn2 = 0, 0
	for _, ch := range nums {
		if ch == ca1 {
			cn1++
		}
		if ch == ca2 {
			cn2++
		}
	}
	if cn1 > len(nums)/3 {
		ans = append(ans, ca1)
	}
	if cn2 > len(nums)/3 && ca2 != ca1 {
		ans = append(ans, ca2)
	}
	return ans
}
