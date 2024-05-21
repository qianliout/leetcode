package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates3([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates4([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	// arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	// bubbleSort(arr)
	// fmt.Println(arr)
}

func findDuplicates(nums []int) []int {
	idx := 0
	for idx < len(nums) {
		d := nums[idx] - 1
		if nums[idx] != nums[d] {
			nums[d], nums[idx] = nums[idx], nums[d]
		} else {
			idx++
		}
	}

	ans := make([]int, 0)
	for i, ch := range nums {
		if i+1 != ch {
			ans = append(ans, ch)
		}
	}
	return ans
}

func findDuplicates2(nums []int) []int {
	ans := make([]int, 0)
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return ans
}

func findDuplicates3(nums []int) (ans []int) {
	for i := range nums {
		for nums[i]-1 != i {
			if nums[i]-1 < 0 || nums[i]-1 >= len(nums) || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return
}

func findDuplicates4(nums []int) []int {
	for i := range nums {
		for nums[i] != i+1 {
			if nums[i]-1 < 0 || nums[i]-1 >= len(nums) || nums[i] == nums[nums[i]-1] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	ans := make([]int, 0)
	for i, ch := range nums {
		if ch != i+1 {
			ans = append(ans, ch)
		}
	}
	return ans
}
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
