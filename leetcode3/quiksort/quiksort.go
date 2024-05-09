package main

func main() {

}

func partition(list []int, low, high int) int {
	pivot := list[low] // 导致 low 位置值为空
	for low < high {
		// high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		// 填补low位置空值
		// high指针值 < pivot high值 移到low位置
		// high 位置值空
		list[low] = list[high]
		// low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		// 填补high位置空值
		// low指针值 > pivot low值 移到high位置
		// low位置值空
		list[high] = list[low]
	}
	// pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		// 位置划分
		pivot := partition(list, low, high)
		// 左边部分排序
		QuickSort(list, low, pivot-1)
		// 右边排序
		QuickSort(list, pivot+1, high)
	}
}
