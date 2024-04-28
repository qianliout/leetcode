package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber2([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber3([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}

func missingNumber2(nums []int) int {
	hasO := false
	for i := range nums {
		if nums[i] == 0 {
			hasO = true
		}
	}
	if !hasO {
		return 0
	}
	nums = append(nums, 0)
	start := 0
	for start < len(nums) {
		ch := nums[start]
		if ch != nums[ch] {
			nums[start], nums[ch] = nums[ch], nums[start]
		} else {
			start++
		}
	}
	for i, ch := range nums {
		if i != ch {
			return i
		}
	}
	return len(nums)
}

func missingNumber(nums []int) int {
	diff := 0
	for i := 0; i <= len(nums); i++ {
		diff = diff ^ i
	}
	for i := 0; i < len(nums); i++ {
		diff = diff ^ nums[i]
	}

	return diff
}

func missingNumber3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i]-1 >= 0 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, ch := range nums {
		if i+1 != ch {
			return i + 1
		}
	}
	return len(nums)
}

func missingNumber4(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for nums[i]-1 >= 0 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, ch := range nums {
		if i+1 != ch {
			return i + 1
		}
	}
	return len(nums)
}
