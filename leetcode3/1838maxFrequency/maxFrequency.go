package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxFrequency([]int{1, 2, 4}, 5))
}

func maxFrequency2(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	// 滑动+二分

	l, r := 0, n
	// 二分，为啥不能l<=r，要好好想想
	for l < r {
		// 这两种写法的区别很重要
		// mid表示区间大小
		// 这种写法，le==ri 时，mid=ri,一定要考虑这一点，不然就会死循环
		mid := (l + r + 1) >> 1
		// fmt.Println("mid", l, r, l+(r+1-l)/2, mid)
		//  找最大,相当于找右边界
		if check(nums, mid, k, sum) {
			l = mid
		} else {
			r = mid - 1
		}
	}

	// 结束时 l=r,所以返回啥都是一样的
	return l
}

func maxFrequency(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	sum := make([]int, n+1)

	for i := 1; i < len(sum); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	// 滑动+二分
	left, right := 0, n+1
	for left < right {
		// 找最大的频率，相当于找右边界
		mid := left + (right-left+1)>>1
		if mid >= 0 && mid < n+1 && check(nums, mid, k, sum) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	// 这个题一定是有答案的，可以不检查最后的结果，但是检测一次是个好习惯
	if left < 0 || left >= n+1 || !check(nums, left, k, sum) {
		return -1
	}

	// 结束时 l=r,所以返回啥都是一样的
	return left
}

func check(nums []int, length, k int, sum []int) bool {
	for i := 0; i+length-1 < len(nums); i++ {
		r := i + length - 1
		// 当前区间合
		cur := sum[r+1] - sum[i]
		// 当前区间的数全部变成最后一个数之后的和
		// 因为 nums 经过了排序，最后一个数是最大值
		t := nums[r] * length

		if t-cur <= k {
			return true
		}
	}
	return false
}
