package main

import (
	"fmt"
)

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p, left, right := m+n-1, m-1, n-1

	for p >= 0 {
		if left >= 0 && right >= 0 && nums1[left] >= nums2[right] {
			nums1[p] = nums1[left]
			p--
			left--
			continue
		}
		if left >= 0 && right >= 0 && nums1[left] < nums2[right] {
			nums1[p] = nums2[right]
			p--
			right--
			continue
		}
		if left >= 0 {
			nums1[p] = nums1[left]
			p--
			left--
			continue
		}
		if right >= 0 {
			nums1[p] = nums2[right]
			p--
			right--
			continue
		}

	}
	fmt.Println(nums1)
}
