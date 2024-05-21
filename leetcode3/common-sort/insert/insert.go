package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 4, 2, 2, 2, 20, 3}
	// insertSort(nums)
	// SelectionSort(nums)
	// BubbleSort(nums)
	ShellSort(nums)
	fmt.Println(nums)

}

// 插入排序，会有元素的移动
func insertSort1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// func insertionSort(data Interface, a, b int) {
// 	for i := a + 1; i < b; i++ {
// 		for j := i; j > a && data.Less(j, j-1); j-- {
// 			data.Swap(j, j-1)
// 		}
// 	}
// }

// 选择排序和插入排序类似，也将数组分为已排序和未排序两个区间。但是在选择排序的实现过程中，不会发生元素的移动，而是直接进行元素的交换。
//
// 选择排序的实现过程: 在不断未排序的区间中找到最小的元素，将其放入已排序区间的尾部。

func SelectionSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		mia := i
		for j := i + 1; j < len(nums); j++ {
			if nums[mia] > nums[j] {
				mia = j
			}
		}
		nums[i], nums[mia] = nums[mia], nums[i]
	}
}

func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

func ShellSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for gap := len(nums) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(nums); i++ {
			tem := nums[i]
			j := i
			for j := i; j >= gap; j = j - gap {
				if nums[j-gap] >= tem {
					break
				}
				nums[j] = nums[j-gap]
			}
			nums[j] = tem
		}
	}
}
