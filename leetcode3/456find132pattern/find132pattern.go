package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(find132pattern([]int{1, 0, 1, -4, -3}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
}

// timeout
func find132pattern2(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				// 找右边
				for k := i + 1; k < len(nums); k++ {
					if nums[k] > nums[j] && nums[k] < nums[i] {
						return true
					}
				}
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/132-pattern/solutions/676970/xiang-xin-ke-xue-xi-lie-xiang-jie-wei-he-95gt/
// 推荐这种写法，更容易理解
func find132pattern(nums []int) bool {
	// 递减栈出栈时最后一次出栈的元素
	k := math.MinInt64
	// 从后到前遍历，单调递减
	stark := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		if len(stark) == 0 {
			stark = append(stark, nums[i])
			continue
		}
		if len(nums)-i >= 2 {
			if nums[i] < k && nums[i] < stark[0] {
				return true
			}
		}
		for len(stark) > 0 && stark[len(stark)-1] < nums[i] {
			// k = Max(k, stark[len(stark)-1])
			k = stark[len(stark)-1]
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, nums[i])
	}
	return false
}

// 是上面方法的简写版本
func find132pattern3(nums []int) bool {
	// 递减栈出栈时最后一次出栈的元素
	k := math.MinInt64
	// 从后到前遍历，单调递减
	stark := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		// k的初值是最小值，如果 k 没有更新过，这里是不能满足的，这是可以简化的关键
		if nums[i] < k {
			return true
		}

		for len(stark) > 0 && stark[len(stark)-1] < nums[i] {
			k = stark[len(stark)-1]
			stark = stark[:len(stark)-1]
		}
		stark = append(stark, nums[i])
	}
	return false
}
