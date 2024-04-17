package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	// fmt.Println(maxNumber([]int{3, 9}, []int{8, 9}, 3))
	// fmt.Println(maxNumber([]int{8, 6, 9}, []int{1, 7, 5}, 3))
	a := []int{2, 5, 6, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0}
	b := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}

	fmt.Println(merge(a, b))
	fmt.Println(merge2(a, b))

	fmt.Println(maxNumber(a, b, 15))
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, 0)
	for i := 0; i <= k; i++ {
		k1, k2 := i, k-i
		if len(nums1) < k1 || len(nums2) < k2 {
			continue
		}
		a1 := pikMax(nums1, k1)
		a2 := pikMax(nums2, k2)
		a3 := merge2(a1, a2)
		if big(a3, ans) {
			ans = a3
		}
	}
	return ans
}

// 找最大的
func pikMax(num []int, k int) []int {
	drop := len(num) - k
	window := make([]int, 0)
	for _, ch := range num {
		for drop > 0 && len(window) > 0 && window[len(window)-1] < ch {
			drop--
			window = window[:len(window)-1]
		}
		window = append(window, ch)
	}

	return window[:k]
}

func merge2(num1, num2 []int) []int {
	ans := make([]int, 0)
	i1, i2 := 0, 0
	for i1 < len(num1) && i2 < len(num2) {
		if big(num1[i1:], num2[i2:]) {
			ans = append(ans, num1[i1])
			i1++
		} else {
			ans = append(ans, num2[i2])
			i2++
		}
	}
	if i1 < len(num1) {
		ans = append(ans, num1[i1:]...)
	}
	if i2 < len(num2) {
		ans = append(ans, num2[i2:]...)
	}
	return ans
}

/*
这种情况下就得不到正确答案
a := []int{2, 5, 6, 4, 4, 0}
b := []int{7, 3, 8, 0, 6, 5, 7, 6, 2}
因为如果一个数相同，则需要比较后面的数据
*/
func merge(num1, num2 []int) []int {
	ans := make([]int, 0)
	i1, i2 := 0, 0
	for i1 < len(num1) && i2 < len(num2) {
		if num1[i1] >= num2[i2] {
			ans = append(ans, num1[i1])
			i1++
		} else {
			ans = append(ans, num2[i2])
			i2++
		}
	}
	if i1 < len(num1) {
		ans = append(ans, num1[i1:]...)
	}
	if i2 < len(num2) {
		ans = append(ans, num2[i2:]...)
	}
	return ans
}

func big(num1, num2 []int) bool {
	i1, i2 := 0, 0
	for i1 < len(num1) && i2 < len(num2) {
		if num1[i1] > num2[i2] {
			return true
		} else if num1[i1] < num2[i2] {
			return false
		} else {
			i1++
			i2++
		}
	}
	return len(num1) > len(num2)
}
