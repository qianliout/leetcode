package main

import (
	"fmt"
)

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 6, 7}))
}

func summaryRanges(nums []int) []string {
	ans := make([]string, 0)
	if len(nums) == 0 {
		return ans
	}
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 < nums[i] {
			if i-1 == start {
				ans = append(ans, fmt.Sprintf("%d", nums[start]))
			} else {
				ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i
		}

	}

	if len(nums)-1 == start {
		ans = append(ans, fmt.Sprintf("%d", nums[start]))
	} else {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	}

	return ans
}
