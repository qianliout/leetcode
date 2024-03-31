package main

import (
	"fmt"
)

func main() {
	ans := []int{2, 0, 2, 1, 1, 0, 1, 2, 1, 1, 2, 0}
	sortColors(ans)
	fmt.Println(ans)
}

func sortColors(nums []int) {
	i, r, w, b := 0, 0, 0, len(nums)-1
	for i <= b {
		if nums[i] == 0 {
			nums[i], nums[r] = nums[r], nums[i]
			r++
			i++
			continue
		} else if nums[i] == 2 {
			nums[i], nums[b] = nums[b], nums[i]
			b--
			continue
		} else if nums[i] == 1 {
			w++
			i++
		}
	}
}
