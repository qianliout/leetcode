package main

import (
	. "outback/geeke/leetcode/common/utils"
)

func main() {

}

func wiggleMaxLength2(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	preDiff := nums[1] - nums[0]
	res := 1
	if preDiff != 0 {
		res = 2
	}
	for i := 2; i < len(nums); i++ {
		curdiff := nums[i] - nums[i-1]
		// 这里的易错点是 prediff<=0 和 prediff>=0
		// 为啥要加上等号呢
		// [3,3,3,2,5]
		// 因为题目要求的是最长摆动子序列的长度，所以只需要统计数组的峰值数量就可以了（相当于是删除单一坡度上的节点，然后统计长度），
		// 这就是贪心所贪的地方，让峰值尽可能的保持峰值，然后删除单一坡度上的节点
		if (curdiff > 0 && preDiff <= 0) || (curdiff < 0 && preDiff >= 0) {
			res++
			preDiff = curdiff
		}
	}
	return res
}

func wiggleMaxLength(nums []int) int {
	// up[i]表示选定idx=i 这个数的最长上升子序列（最后一个向上走）
	up := make([]int, len(nums))
	// up[i]表示选定idx=i 这个数的最长下绛子序列（最后一个向下走）
	down := make([]int, len(nums))
	for i := range nums {
		up[i], down[i] = 1, 1
	}
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				// 从j->i 只能是向上走
				up[i] = Max(up[i], down[j]+1)
			}
			if nums[i] < nums[j] {
				// 从j->i 只能是向下走
				down[i] = Max(down[i], up[j]+1)
			}
		}
	}
	return Max(up[len(nums)-1], down[len(nums)-1])
}
