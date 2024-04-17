package main

import (
	"fmt"
	"sort"
)

func main() {
	// wiggleSort([]int{1, 3, 2, 2, 3, 1})
	wiggleSort([]int{1, 4, 3, 4, 1, 2, 1, 3, 1, 3, 2, 3, 3})
}

func wiggleSort(nums []int) {
	tem := make([]int, len(nums))
	copy(tem, nums)

	sort.Ints(tem)
	mid := (len(tem) + 1) / 2
	left := tem[:mid]
	right := tem[mid:]
	reverse(left)
	reverse(right)
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		nums[i] = left[l]
		l++
		i++
		nums[i] = right[r]
		r++
		i++
	}
	// left 可能比 right 多一个
	if l < len(left) {
		nums[i] = left[r]
	}
	fmt.Println(nums)
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
