package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 4, 2, 2, 2, 20, 3}
	MergeSort(nums)
	// SelectionSort(nums)
	// BubbleSort(nums)
	fmt.Println(nums)
}

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// if len(nums) == 2 {
	// 	if nums[0] > nums[1] {
	// 		nums[1], nums[0] = nums[0], nums[1]
	// 	}
	// 	return
	// }

	mid := len(nums) / 2
	MergeSort(nums[:mid])
	MergeSort(nums[mid:])
	Merge(nums, mid)
}

// 原地排序，可以使用插入排序的思想
func Merge(nums []int, mid int) {
	for i := mid; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				continue
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
