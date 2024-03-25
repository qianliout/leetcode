package main

import (
	"fmt"
	"sort"
)

func main() {
	num1 := []int{1, 2}
	num2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays1(num1, num2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	return 0
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	if len(nums1) == 0 {
		return 0
	}
	sort.Ints(nums1)

	if len(nums1)%2 == 0 {
		a := (len(nums1) - 1) / 2
		b := a + 1
		return (float64(nums1[a]) + float64(nums1[b])) / 2
	}

	a := (len(nums1)) / 2
	return float64(nums1[a])
}
