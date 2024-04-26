package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(splitArray([]int{2, 16, 14, 15}, 2))
	fmt.Println(split([]int{2, 14, 15, 16}, 29))
}

/*
给定一个非负整数数组 nums 和一个整数 k ，你需要将这个数组分成 k 个非空的连续子数组。
设计一个算法使得这 k 个子数组各自和的最大值最小。
*/
func splitArray(nums []int, k int) int {
	sum, ma := 0, math.MinInt64
	for _, ch := range nums {
		sum += ch
		if ch > ma {
			ma = ch
		}
	}
	if k >= len(nums) {
		return ma
	}
	// if k <= 1 {
	// 	return sum
	// }

	left, right := ma, sum
	for left <= right {
		mid := left + (right-left+1)/2
		m := split(nums, mid)
		if m == k {
			// 看一下是否还有更小的
			right = mid - 1
		} else if m > k {
			// 说明，还可以放更大的值
			left = mid + 1
		} else if m < k {
			// 说明,mid 放进去之后分的次数少，说明还可以放一个更多小的值
			right = mid - 1
		}
	}
	// 检测边界
	if left < 0 || split(nums, left) > k || left > sum {
		return -1
	}

	return left
}

// 各组和和最大值是 maxSum,最少要拆分成多少组
func split(nums []int, maxSum int) int {
	ans := 0
	sum := 0
	for _, ch := range nums {
		if sum+ch > maxSum {
			ans++
			sum = 0
		}
		sum += ch
	}
	return ans + 1
}
