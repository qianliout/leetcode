package main

import (
	"math"

	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

/*
给定一个非负整数数组 nums 和一个整数 k ，你需要将这个数组分成 k 个非空的连续子数组。
设计一个算法使得这 k 个子数组各自和的最大值最小。
*/
func splitArray(nums []int, k int) int {
	sum, ma := 0, math.MinInt64
	for _, ch := range nums {
		sum += ch
		ma = Max(ma, ch)
	}
	le, ri := ma, sum
	for le < ri {
		mid := le + (ri-le)/2
		m := split(nums, mid)
		if m == k {
			ri = mid
		} else if m > k {
			ri = mid - 1
		} else if m < k {
			le = mid + 1
		}
	}
	return le
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

	return ans
}
