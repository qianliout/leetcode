package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 4, 2, 2, 2, 20, 3}
	// insertSort(nums)
	// SelectionSort(nums)
	// BubbleSort(nums)
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func partition(nums []int, low, high int) int {
	pivot := nums[low]
	// 以nums[low]做为 pivot 元素
	// j 表示小堆的最后一个元素的位置，当还没有小堆时，就是low
	j := low
	// i 表示乱堆的开始位置
	for i := low + 1; i <= high; i++ {
		ch := nums[i]
		// 如果ch 大于 pivot 说明他本来就应该在大堆中，所以只需要移动 i就行（i++的操作）
		// if ch >= pivot {
		// 	continue
		// }

		// 说明在小堆中，需要做两步
		// 1：把 ch 和大堆中的第一个元素交换，
		// 2：把 j 的位置向后移动一位
		if ch < pivot {
			j++
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	// 都做完之后，把第一个元素(pivot)放到中间，这里可以取个巧，把小堆的最后一个元素和 pivot 交换就好
	nums[low], nums[j] = nums[j], nums[low]
	return j
}

func QuickSort(nums []int, low, high int) {
	if len(nums) <= 1 || low >= high || len(nums)-1 < low {
		return
	}
	// 位置划分
	pivot := partition(nums, low, high)
	// 左边部分排序
	QuickSort(nums, low, pivot-1)
	// 右边排序
	QuickSort(nums, pivot+1, high)
}
